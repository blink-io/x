package noop

import "google.golang.org/grpc/grpclog"

var _ grpclog.LoggerV2 = (*logger)(nil)

type logger struct{}

func (l *logger) Info(args ...any) {

}

func (l *logger) Infoln(args ...any) {

}

func (l *logger) Infof(format string, args ...any) {

}

func (l *logger) Warning(args ...any) {

}

func (l *logger) Warningln(args ...any) {

}

func (l *logger) Warningf(format string, args ...any) {

}

func (l *logger) Error(args ...any) {

}

func (l *logger) Errorln(args ...any) {

}

func (l *logger) Errorf(format string, args ...any) {

}

func (l *logger) Fatal(args ...any) {

}

func (l *logger) Fatalln(args ...any) {

}

func (l *logger) Fatalf(format string, args ...any) {

}

func (l *logger) V(lv int) bool {
	return false
}
