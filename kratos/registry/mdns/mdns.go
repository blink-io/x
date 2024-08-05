package mdns

import (
	"context"
	"net"
	"time"

	kregistry "github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/mdns"
)

var (
	_ kregistry.Registrar = (*Registry)(nil)
	_ kregistry.Discovery = (*Registry)(nil)
)

type mdnsEntry struct {
	node *mdns.Server
	id   string
}

type Registry struct {
	services map[string][]*mdnsEntry

	instances map[string]*kregistry.ServiceInstance
	opts      *options
}

func New(opts ...Option) (r *Registry) {
	op := &options{
		ctx:       context.Background(),
		namespace: "/microservices",
		ttl:       time.Second * 15,
		maxRetry:  5,
	}
	for _, o := range opts {
		o(op)
	}
	return &Registry{
		opts: op,
	}
}

func (r *Registry) GetService(ctx context.Context, serviceName string) ([]*kregistry.ServiceInstance, error) {
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (kregistry.Watcher, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Register(ctx context.Context, service *kregistry.ServiceInstance) error {
	domain := ""
	hostName := ""
	port := 9999
	msvc, err := mdns.NewMDNSService(
		service.ID,
		service.Name,
		domain,
		hostName,
		port,
		[]net.IP{net.ParseIP("0.0.0.0")},
		nil,
	)
	if err != nil {
		return err
	}

	srv, err := mdns.NewServer(&mdns.Config{Zone: msvc})
	if err != nil {
		return err
	}

	r.services[service.Name] = []*mdnsEntry{
		{
			node: srv,
			id:   service.ID,
		},
	}
	return nil
}

func (r *Registry) Deregister(ctx context.Context, service *kregistry.ServiceInstance) error {
	//TODO implement me
	panic("implement me")
}
