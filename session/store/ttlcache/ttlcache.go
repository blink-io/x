package ttlcache

import (
	"context"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"

	"github.com/blink-io/x/session/store"
)

const Name = "ttlcache"

var _ store.Store = (*Store)(nil)

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() *Store {
	cc := ttlcache.New[string, []byte]()
	// Without Start the cache never deletes expired items on its own and
	// grows unboundedly.
	go cc.Start()

	return &Store{cc: cc, tt: store.NewTokenMap()}
}

var _ store.Store = (*Store)(nil)

type Store struct {
	// mu guards tt; cc is safe for concurrent use on its own.
	mu sync.RWMutex
	tt store.TokenMap
	cc *ttlcache.Cache[string, []byte]
}

func (s *Store) Name() string {
	return Name
}

func (s *Store) Delete(ctx context.Context, token string) (err error) {
	s.cc.Delete(token)
	s.mu.Lock()
	delete(s.tt, token)
	s.mu.Unlock()
	return nil
}

func (s *Store) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if item := s.cc.Get(token); item != nil && !item.IsExpired() {
		return item.Value(), true, nil
	}
	return nil, false, nil
}

func (s *Store) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	ttl := time.Until(expiry)
	s.cc.Set(token, data, ttl)
	s.mu.Lock()
	s.tt[token] = store.NilStruct
	s.mu.Unlock()
	return nil
}

func (s *Store) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for token := range s.tt {
		item := s.cc.Get(token)
		if item == nil || item.IsExpired() {
			// Prune tokens whose data has expired; delete during range is
			// safe in Go and keeps the token map from growing forever.
			delete(s.tt, token)
			continue
		}
		sessions[token] = item.Value()
	}
	return sessions, nil
}
