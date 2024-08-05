package loader

import (
	"io"
	"net/http"

	"github.com/blink-io/x/i18n"
)

// loader loads by http GET requests
// URLVar should be like: https://xxx.com/languages/zh_Hans.toml
type loader struct {
	url  string
	opts *options
}

func NewLoader(url string, ops ...Option) i18n.Loader {
	return newLoader(url, ops...)
}

func newLoader(url string, ops ...Option) *loader {
	opts := applyOptions(ops...)
	return &loader{url: url, opts: opts}
}

func LoadFromHTTP(url string, ops ...Option) error {
	ld := newLoader(url, ops...)
	return ld.Load(ld.opts.bundle)
}

func (l *loader) Load(b i18n.Bundler) error {
	c := l.opts.client

	var res *http.Response
	var err error
	if rf := l.opts.requestFunc; rf != nil {
		if req, rerr := l.opts.requestFunc(c, l.url); rerr != nil {
			return rerr
		} else {
			res, err = c.Do(req)
		}
	} else {
		res, err = c.Get(l.url)
	}
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var path string
	if ef := l.opts.extractFunc; ef != nil {
		path = ef(l.url)
	} else {
		path = l.url
	}
	if _, err := b.LoadMessageFileBytes(buf, path); err != nil {
		i18n.GetLogger()("[WARN] unable to load message from URL: %s", l.url)
	}
	return nil
}
