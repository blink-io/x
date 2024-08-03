package roller

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/natefinch/lumberjack/v3"
)

const (
	defaultLogMaxSize = 300 // MB
)

// Option serializes file log related config in toml/json.
type Option struct {
	Level slog.Level

	// Log filename, leave empty to disable file log.
	Filename string `toml:"filename" json:"filename"`
	// Max size for a single file, in MB.
	MaxSize int64 `toml:"max-size" json:"max-size"`
	// MaxAge is the maximum time to retain old log files based on the timestamp
	// encoded in their filename. The default is not to remove old log files
	// based on age.
	MaxAge time.Duration `toml:"max-age" json:"max-age"`
	// Maximum number of old log files to retain.
	MaxBackups int `toml:"max-backups" json:"max-backups"`
	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time. The default is to use UTC
	// time.
	LocalTime bool `toml:"local-time" json:"local-time"`
}
type Handler struct {
	roller *lumberjack.Roller
	opt    *Option
}

func NewHandler(o *Option) (*Handler, error) {
	roller, err := initFileRoller(o)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		roller: roller,
	}
	return h, nil
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opt.Level
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) WithGroup(name string) slog.Handler {
	//TODO implement me
	panic("implement me")
}

var _ slog.Handler = &Handler{}

// initFileRoller initializes file based logging options.
func initFileRoller(o *Option) (*lumberjack.Roller, error) {
	if st, err := os.Stat(o.Filename); err == nil {
		if st.IsDir() {
			return nil, errors.New("can't use directory as log file name")
		}
	}
	if o.MaxSize == 0 {
		o.MaxSize = defaultLogMaxSize
	}

	rops := &lumberjack.Options{
		MaxAge:     o.MaxAge,
		MaxBackups: o.MaxBackups,
		LocalTime:  o.LocalTime,
	}
	return lumberjack.NewRoller(o.Filename, o.MaxSize, rops)
}
