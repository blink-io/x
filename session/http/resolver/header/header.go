package header

import (
	"net/http"
	"time"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
	. "github.com/blink-io/x/session/http/shared"
)

const DefaultHeader = "X-Auth-Token"

var _ resolver.Resolver = (*rv)(nil)

type rv struct {
	header       string
	exposeExpiry bool
}

func Default() resolver.Resolver {
	return New(DefaultHeader)
}

func New(h string) resolver.Resolver {
	return &rv{
		header: h,
	}
}

func (v *rv) Resolve(m resolver.Manager, ef resolver.ErrorFunc, w http.ResponseWriter, r *http.Request, next http.Handler) error {
	token := r.Header.Get(v.header)

	ctx, err := m.Load(r.Context(), token)
	if err != nil {
		return err
	}

	sr := r.WithContext(ctx)

	sw := &SessionResponseWriter{
		CommitAndWriteSession: func(w http.ResponseWriter, r *http.Request) {
			v.commitAndWriteSessionHeader(m, ef, w, sr)
		},
		ResponseWriter: w,
		Request:        sr,
	}

	next.ServeHTTP(sw, sr)

	if !sw.IsWritten() {
		v.commitAndWriteSessionHeader(m, ef, w, sr)
	}
	return nil
}

func (v *rv) commitAndWriteSessionHeader(m resolver.Manager, ef resolver.ErrorFunc, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch m.Status(ctx) {
	case session.Modified:
		token, expiry, err := m.Commit(ctx)
		if err != nil {
			ef(w, r, err)
			return
		}

		w.Header().Add(v.header, token)
		if v.exposeExpiry {
			w.Header().Add(v.header+"-Expires", expiry.Format(time.RFC3339Nano))
		}
	case session.Destroyed:
		w.Header().Del(v.header)
	}
}
