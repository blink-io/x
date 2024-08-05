package realip

import (
	"net/http"

	"github.com/blink-io/x/realip"
)

type Options = realip.Options

func New(ops ...Option) func(http.Handler) http.Handler {
	o := applyOption(ops...)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ip := getFromHTTP(r, o); len(ip) > 0 {
				r = r.WithContext(realip.NewContext(r.Context(), ip))
			}
			h.ServeHTTP(w, r)
		})
	}
}
