package loader

import (
	"context"
	"time"

	"github.com/blink-io/x/i18n"

	"google.golang.org/grpc"
)

// loader loads by GRPC services
type loader struct {
	client    I18NClient
	endpoint  string
	languages []string
}

func NewLoader(cc grpc.ClientConnInterface, languages []string) i18n.Loader {
	return newLoader(cc, languages)
}

func newLoader(cc grpc.ClientConnInterface, languages []string) *loader {
	client := NewI18NClient(cc)
	return &loader{client: client, languages: languages}
}

func LoadFromGRPC(cc grpc.ClientConnInterface, languages []string, ops ...Option) error {
	opts := applyOptions(ops...)
	return NewLoader(cc, languages).Load(opts.bundle)
}

func (l *loader) Load(b i18n.Bundler) error {
	req := &ListLanguagesRequest{
		Languages: l.languages,
	}
	res, err := l.client.ListLanguages(context.Background(), req)
	if err != nil {
		return err
	}
	for _, v := range res.Entries {
		// Ignore invalid
		if !v.Valid {
			continue
		}
		if _, verr := b.LoadMessageFileBytes(v.Payload, v.Path); verr != nil {
			i18n.GetLogger()("[WARN] unable to load message from gRPC service: %s, endpoint: %s, reason: %s", v.Path, l.endpoint, verr.Error())
		}
	}
	return nil
}

type grpcServer struct {
	UnimplementedI18NServer
	h i18n.EntryHandler
}

func newGRPCServer(h i18n.EntryHandler) *grpcServer {
	gsrv := &grpcServer{h: h}
	return gsrv
}

func (s *grpcServer) ListLanguages(ctx context.Context, req *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := &LanguageEntry{
				Language: l,
			}
			if e := em[l]; e != nil {
				le.Path = e.Path
				le.Payload = e.Payload
				le.Valid = true
			} else {
				le.Valid = false
			}
			entries[l] = le
		}
	}

	res := &ListLanguagesResponse{
		Entries:   entries,
		Timestamp: time.Now().Unix(),
	}

	return res, nil
}

func RegisterEntryHandler(gsrv *grpc.Server, eh i18n.EntryHandler) {
	ss := newGRPCServer(eh)
	RegisterI18NServer(gsrv, ss)
}

func RegisterEntryHandlerFunc(gsrv *grpc.Server, ehf i18n.EntryHandlerFunc) {
	ss := newGRPCServer(ehf)
	RegisterI18NServer(gsrv, ss)
}
