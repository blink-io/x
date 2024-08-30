package orm

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestDBR_1(t *testing.T) {
	cc, err := dbr.Open("sqlite", "./orm_demo.db", &loggerEventReceiver{})
	require.NoError(t, err)

	m := new(Model)

	r := cc.QueryRow("select 'heison' as name, sqlite_version() as version")
	require.NoError(t, r.Scan(&m.Name, &m.Version))

	fmt.Println(m)
}

type loggerEventReceiver struct {
	sl *slog.Logger
}

func (l *loggerEventReceiver) Event(eventName string) {
	l.sl.Info("", "event", eventName)
}

func (l *loggerEventReceiver) EventKv(eventName string, kvs map[string]string) {
	// attrs := make([]slog.Attr, 0, len(kvs))
}

func (l *loggerEventReceiver) EventErr(eventName string, err error) error {

	return nil
}

func (l *loggerEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {

	return nil
}

func (l *loggerEventReceiver) Timing(eventName string, nanoseconds int64) {

}

func (l *loggerEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {

}

var _ dbr.EventReceiver = (*loggerEventReceiver)(nil)
