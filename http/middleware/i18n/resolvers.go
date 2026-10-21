package i18n

import (
	stdhttp "net/http"
	"strings"
	"sync"
	"sync/atomic"

	"golang.org/x/text/language"
)

// Resolver gets locale info from HTTP/gRPC transport
type Resolver interface {
	Resolve(*stdhttp.Request) string
}

var (
	_ Resolver = (*HeaderResolver)(nil)
	_ Resolver = (*CookieResolver)(nil)
	_ Resolver = (*QueryParamsResolver)(nil)
	_ Resolver = (*AcceptLanguageResolver)(nil)
	_ Resolver = (*PathResolver)(nil)
)

// HeaderResolver gets language from header in HTTP
type HeaderResolver struct {
	Name string
}

func (rv *HeaderResolver) Resolve(r *stdhttp.Request) string {
	val := r.Header.Get(rv.Name)
	return val
}

// CookieResolver gets locale info from HTTP cookie
type CookieResolver struct {
	Name string
}

func (rv *CookieResolver) Resolve(r *stdhttp.Request) string {
	if c, err := r.Cookie(rv.Name); err == nil {
		return c.Value
	}
	return ""
}

// QueryParamsResolver gets locale info from query string
type QueryParamsResolver struct {
	Name string
}

func NewQueryParamsResolver(name string) Resolver {
	return &QueryParamsResolver{
		Name: name,
	}
}

func (rv *QueryParamsResolver) Resolve(r *stdhttp.Request) string {
	// r.URL.Query() parses the entire RawQuery and allocates two maps; a
	// single check avoids all of it when the parameter is absent, which is
	// the common case.
	if !strings.Contains(r.URL.RawQuery, rv.Name+"=") {
		return ""
	}
	return r.URL.Query().Get(rv.Name)
}

// AcceptLanguageResolver gets locale info from header accept-language
// example of accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7
type AcceptLanguageResolver struct{}

func NewAcceptLanguageResolver() Resolver {
	return &AcceptLanguageResolver{}
}

func (rv *AcceptLanguageResolver) Resolve(r *stdhttp.Request) string {
	val := r.Header.Get("Accept-Language")
	if val == "" {
		return ""
	}
	// The Accept-Language header repeats heavily in real traffic, and
	// language.ParseAcceptLanguage is not cheap; memoize by the raw header
	// value up to a bounded number of distinct values.
	if v, ok := acceptLangCache.Load(val); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	tags, _, err := language.ParseAcceptLanguage(val)
	if err != nil || len(tags) == 0 {
		return ""
	}
	// ParseAcceptLanguage orders tags by quality; the first entry is the
	// preferred language.
	tag := tags[0].String()
	if acceptLangCount.Add(1) > acceptLangCacheMax {
		return tag
	}
	acceptLangCache.Store(val, tag)
	return tag
}

var (
	// Both are shared between resolver instances; the count bounds memory
	// in case of adversarial header values.
	acceptLangCache sync.Map
	acceptLangCount atomic.Int64
)

const acceptLangCacheMax = 128

// PathResolver gets locale info from url path expr.
// Foe example, /products/{lang} -> /products/zh-cn?product_id=XXX,
type PathResolver struct {
	Key  string
	Expr string
}

func (rv *PathResolver) Resolve(r *stdhttp.Request) string {
	return ""
}
