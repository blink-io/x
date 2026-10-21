package ristretto

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto/v2"

	"github.com/blink-io/x/session/store"
)

const Name = "ristretto"

var _ store.Store = (*Store)(nil)

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() *Store {
	cc, err := ristretto.NewCache[string, []byte](&ristretto.Config[string, []byte]{
		NumCounters: 100_000,
		MaxCost:     100_000,
		BufferItems: 100_000,
	})
	if err != nil {
		log.Printf("unable to create ristretto cache: %v", err)
	}
	s := &Store{cc: cc, tm: store.NewTokenMap()}
	return s
}

var _ store.Store = (*Store)(nil)

type Store struct {
	// mu guards tm; cc is safe for concurrent use on its own.
	mu sync.RWMutex
	tm store.TokenMap
	cc *ristretto.Cache[string, []byte]
}

func (s *Store) Name() string {
	return Name
}

func (s *Store) Delete(ctx context.Context, token string) (err error) {
	s.cc.Del(token)
	s.mu.Lock()
	delete(s.tm, token)
	s.mu.Unlock()
	return nil
}

func (s *Store) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if v, ok := s.cc.Get(token); ok && v != nil {
		return v, true, nil
	}
	return nil, false, nil
}

func (s *Store) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	ttl := time.Until(expiry)
	ok := s.cc.SetWithTTL(token, data, 1, ttl)
	if !ok {
		return fmt.Errorf("unable to store: %s", token)
	}
	s.mu.Lock()
	s.tm[token] = store.NilStruct
	s.mu.Unlock()
	return nil
}

func (s *Store) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for token := range s.tm {
		if data, ok := s.cc.Get(token); ok && data != nil {
			sessions[token] = data
			continue
		}
		// Prune tokens whose data has expired; delete during range is
		// safe in Go and keeps the token map from growing forever.
		delete(s.tm, token)
	}
	return sessions, nil
}
