package slog

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc/grpclog"
)

// See https://github.com/grpc/grpc-go/blob/v1.35.0/grpclog/loggerv2.go#L77-L86
const (
	grpcLvlInfo int = iota
	grpcLvlWarn
	grpcLvlError
	grpcLvlFatal
	grpcLvlDebug

	LevelFatal slog.Level = 12
)

var (
	// _grpcToSlogLevel maps gRPC log levels to slog log levels.
	_grpcToSlogLevel = map[int]slog.Level{
		grpcLvlDebug: slog.LevelDebug,
		grpcLvlInfo:  slog.LevelInfo,
		grpcLvlWarn:  slog.LevelWarn,
		grpcLvlError: slog.LevelError,
		grpcLvlFatal: LevelFatal,
	}
)

type LevelEnabler interface {
	Enabled(ctx context.Context, level slog.Level) bool
}

// An Option overrides a Logger's default configuration.
type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

// withWarn redirects the fatal level to the warn level, which makes testing
// easier. This is intentionally unexported.
func withWarn() Option {
	return optionFunc(func(logger *Logger) {
		logger.fatal = &printer{
			enabler: logger.delegate,
			level:   slog.LevelWarn,
			print:   printWrapper(logger.delegate.Warn),
			printf:  printfWrapper(logger.delegate.Warn),
		}
	})
}

// NewLogger returns a new Logger.
func NewLogger(l *slog.Logger, options ...Option) *Logger {
	logger := &Logger{
		delegate: newLogger(l),
		enabler:  l,
		ctx:      context.Background(),
	}
	logger.print = &printer{
		enabler: logger.enabler,
		level:   slog.LevelInfo,
		print:   printWrapper(logger.delegate.Info),
		printf:  printfWrapper(logger.delegate.Info),
	}
	logger.fatal = &printer{
		enabler: logger.enabler,
		level:   LevelFatal,
		print:   printWrapper(logger.delegate.Fatal),
		printf:  printfWrapper(logger.delegate.Fatal),
	}
	for _, option := range options {
		option.apply(logger)
	}
	return logger
}

// printer implements Print, Printf, and Println operations for a slog level.
//
// We use it to customize Debug vs Info, and Warn vs Fatal for Print and Fatal
// respectively.
type printer struct {
	enabler LevelEnabler
	level   slog.Level
	print   func(...any)
	printf  func(string, ...any)
}

func (v *printer) Print(args ...any) {
	v.print(args...)
}

func (v *printer) Printf(format string, args ...any) {
	v.printf(format, args...)
}

func (v *printer) Println(args ...any) {
	if v.enabler.Enabled(context.Background(), v.level) {
		v.print(sprintln(args))
	}
}

var _ grpclog.LoggerV2 = (*Logger)(nil)

// Logger adapts slog's Logger to be compatible with grpclog.LoggerV2 only.
type Logger struct {
	delegate *logger
	enabler  LevelEnabler
	print    *printer
	fatal    *printer
	ctx      context.Context
	// printToDebug bool
	// fatalToWarn  bool
}

// Info implements grpclog.LoggerV2.
func (l *Logger) Info(args ...any) {
	l.delegate.InfoContext(l.ctx, sprintln(args))
}

// Infoln implements grpclog.LoggerV2.
func (l *Logger) Infoln(args ...any) {
	if l.enabler.Enabled(l.ctx, slog.LevelInfo) {
		l.delegate.InfoContext(l.ctx, sprintln(args))
	}
}

// Infof implements grpclog.LoggerV2.
func (l *Logger) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.delegate.InfoContext(l.ctx, msg)
}

// Warning implements grpclog.LoggerV2.
func (l *Logger) Warning(args ...any) {
	l.delegate.WarnContext(l.ctx, sprintln(args))
}

// Warningln implements grpclog.LoggerV2.
func (l *Logger) Warningln(args ...any) {
	if l.enabler.Enabled(l.ctx, slog.LevelWarn) {
		l.delegate.WarnContext(l.ctx, sprintln(args))
	}
}

// Warningf implements grpclog.LoggerV2.
func (l *Logger) Warningf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.delegate.WarnContext(l.ctx, msg)
}

// Error implements grpclog.LoggerV2.
func (l *Logger) Error(args ...any) {
	l.delegate.ErrorContext(l.ctx, sprintln(args))
}

// Errorln implements grpclog.LoggerV2.
func (l *Logger) Errorln(args ...any) {
	if l.enabler.Enabled(l.ctx, slog.LevelError) {
		l.delegate.ErrorContext(l.ctx, sprintln(args))
	}
}

// Errorf implements grpclog.LoggerV2.
func (l *Logger) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.delegate.ErrorContext(l.ctx, msg)
}

// Fatal implements grpclog.LoggerV2.
func (l *Logger) Fatal(args ...any) {
	l.fatal.Print(args...)
}

// Fatalln implements grpclog.LoggerV2.
func (l *Logger) Fatalln(args ...any) {
	l.fatal.Println(args...)
}

// Fatalf implements grpclog.LoggerV2.
func (l *Logger) Fatalf(format string, args ...any) {
	l.fatal.Printf(format, args...)
}

// V implements grpclog.LoggerV2.
func (l *Logger) V(level int) bool {
	return l.enabler.Enabled(l.ctx, _grpcToSlogLevel[level])
}

func sprintln(args []any) string {
	s := fmt.Sprintln(args...)
	// Drop the new line character added by Sprintln
	return s[:len(s)-1]
}

func printWrapper(f func(string, ...any)) func(...any) {
	return func(args ...any) {
		f(sprintln(args))
	}
}

func printfWrapper(f func(string, ...any)) func(string, ...any) {
	return func(format string, args ...any) {
		msg := fmt.Sprintf(format, args...)
		f(msg)
	}
}

type logger struct {
	*slog.Logger
}

func newLogger(l *slog.Logger) *logger {
	return &logger{l}
}

func (l *logger) Fatal(msg string, args ...any) {
	l.FatalContext(context.Background(), msg, args...)
}

func (l *logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelFatal, msg, args...)
}
