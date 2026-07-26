package registrar

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type RegisterFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	NATSRegistrar(context.Context) RegisterFunc
}

type ServiceRegistrar interface {
	RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error)

	RequestWithContext(ctx context.Context, subj string, data []byte) (*nats.Msg, error)

	JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error)

	Publish(subj string, data []byte) error

	PublishMsg(m *nats.Msg) error

	PublishRequest(subj, reply string, data []byte) error

	RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error)

	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)

	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)

	ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error)

	ChanQueueSubscribe(subj, group string, ch chan *nats.Msg) (*nats.Subscription, error)

	SubscribeSync(subj string) (*nats.Subscription, error)

	QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error)

	QueueSubscribeSync(subj, queue string) (*nats.Subscription, error)

	QueueSubscribeSyncWithChan(subj, queue string, ch chan *nats.Msg) (*nats.Subscription, error)
}

var _ ServiceRegistrar = (*nats.Conn)(nil)

type serviceRegistrar struct {
	*nats.Conn
}

func NewServiceRegistrar(c *nats.Conn) ServiceRegistrar {
	return serviceRegistrar{c}
}

type Func[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToNATS(context.Context, ServiceRegistrar) error
}

type registrar[S any] struct {
	s S
	f Func[S]
}

var _ Registrar = (*registrar[any])(nil)

func (h *registrar[S]) RegisterToNATS(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}

func NewRegistrar[S any](s S, f Func[S]) Registrar {
	rr := &registrar[S]{
		s: s,
		f: f,
	}
	return rr
}
