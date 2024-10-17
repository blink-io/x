package registrar

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type (
	Task          = gocron.Task
	Job           = gocron.Job
	JobOption     = gocron.JobOption
	JobDefinition = gocron.JobDefinition
)

type RegisterFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	GocronRegistrar(context.Context) RegisterFunc
}

type ServiceRegistrar interface {
	// NewJob creates a new job in the Scheduler. The job is scheduled per the provided
	// definition when the Scheduler is started. If the Scheduler is already running
	// the job will be scheduled when the Scheduler is started.
	NewJob(JobDefinition, Task, ...JobOption) (Job, error)
	// RemoveByTags removes all jobs that have at least one of the provided tags.
	RemoveByTags(...string)
	// RemoveJob removes the job with the provided id.
	RemoveJob(uuid.UUID) error
	// Update replaces the existing Job's JobDefinition with the provided
	// JobDefinition. The Job's Job.UserID() remains the same.
	Update(uuid.UUID, JobDefinition, Task, ...JobOption) (Job, error)
}

type serviceRegistrar struct {
	gocron.Scheduler
}

func NewServiceRegistrar(s gocron.Scheduler) ServiceRegistrar {
	return serviceRegistrar{s}
}

type Func[S any] func(ServiceRegistrar, S)

type FuncWithErr[S any] func(ServiceRegistrar, S) error

type CtxFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxFuncWithErr[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToGocron(context.Context, ServiceRegistrar) error
}

var _ Registrar = (*registrar[any])(nil)

func New[S any](s S, f Func[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewWithErr[S any](s S, f FuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxWithErr(s, cf)
}

func NewCtx[S any](s S, f CtxFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtxWithErr[S any](s S, f CtxFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

type registrar[S any] struct {
	s S
	f CtxFuncWithErr[S]
}

func (h *registrar[S]) RegisterToGocron(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
