package requestid

import (
	"net/http"

	"github.com/blink-io/x/requestid"
)

func New(ops ...Option) func(http.Handler) http.Handler {
	opts := applyOptions(ops...)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get(opts.header)
			if len(reqID) == 0 {
				reqID = opts.generator()
			}
			r = r.WithContext(requestid.NewContext(r.Context(), reqID))
			w.Header().Set(opts.header, reqID)
			h.ServeHTTP(w, r)
		})
	}
}
