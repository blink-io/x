package registrar

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	qslog "github.com/blink-io/x/scheduler/quartz/logger/slog"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/logger"
	"github.com/reugn/go-quartz/quartz"
	"github.com/stretchr/testify/require"
)

type svc struct{}

func (s *svc) QuartzRegistrar(ctx context.Context) RegisterFunc {
	return func(ctx context.Context, rr ServiceRegistrar) error {
		return s.doRegister(rr)
	}
}

func (s *svc) doRegister(rr ServiceRegistrar) error {
	fnJob := job.NewFunctionJob[string](func(ctx context.Context) (string, error) {
		fmt.Println("Hello, Quartz")
		return "Hello, Quartz", nil
	})
	err := rr.ScheduleJob(
		quartz.NewJobDetail(fnJob, quartz.NewJobKey("functionJob")),
		quartz.NewSimpleTrigger(time.Second*5),
	)
	return err
}

var _ WithRegistrar = (*svc)(nil)

func TestIface(t *testing.T) {
	logger.SetDefault(qslog.New(slog.Default()))

	ctx := context.Background()
	sched := quartz.NewStdScheduler()

	rr := NewServiceRegistrar(sched)
	require.NotNil(t, rr)

	var s = &svc{}

	err := s.QuartzRegistrar(ctx)(ctx, rr)
	require.NoError(t, err)

	sched.Start(ctx)

	fmt.Println("Quartz .....")

	time.Sleep(50 * time.Second)
}
