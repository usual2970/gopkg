package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/usual2970/gopkg/runenv"
	"github.com/zeromicro/go-zero/core/logx"
)

type DLogger struct {
	logger logx.Logger
	fields []logx.LogField
}

var (
	dlog *DLogger

	Debug, Error, Info, Slow     func(...interface{})
	Debugf, Errorf, Infof, Slowf func(string, ...interface{})
	Debugv, Errorv, Infov, Slowv func(interface{})
	Debugw, Errorw, Infow, Sloww func(string, ...logx.LogField)

	WithCallerSkip func(skip int) *DLogger

	WithContext func(ctx context.Context) *DLogger

	WithDuration func(d time.Duration) *DLogger

	WithFields func(map[string]interface{}) *DLogger
	WithField  func(string, interface{}) *DLogger
)

const (
	serviceNameKey     = "HUB_SERVICE"
	defaultServiceName = "hub"
)

func Setup() {
	conf := logx.LogConf{
		ServiceName: getServiceName(),
		Mode:        getMode(),
	}
	logx.MustSetup(conf)
	dlog = &DLogger{
		logger: logx.WithCallerSkip(2),
		fields: make([]logx.LogField, 0),
	}

	AddGlobalFields(map[string]interface{}{
		"service": getServiceName(),
	})

	Debug, Error, Info, Slow = dlog.Debug, dlog.Error, dlog.Info, dlog.Slow
	Debugf, Errorf, Infof, Slowf = dlog.Debugf, dlog.Errorf, dlog.Infof, dlog.Slowf
	Debugv, Errorv, Infov, Slowv = dlog.Debugv, dlog.Errorv, dlog.Infov, dlog.Slowv
	Debugw, Errorw, Infow, Sloww = dlog.Debugw, dlog.Errorw, dlog.Infow, dlog.Sloww

	WithFields = dlog.WithFields
	WithField = dlog.WithField
}

func getServiceName() string {
	rs := os.Getenv(serviceNameKey)
	if rs == "" {
		rs = defaultServiceName
	}
	return rs
}

func getMode() string {
	if runenv.IsDev() {
		return "console"
	}
	return "file"
}

type (
	LogField = logx.LogField
)

func AddGlobalFields(fields map[string]interface{}) {
	rs := make([]LogField, 0, len(fields))
	for k, v := range fields {
		rs = append(rs, logx.Field(k, v))
	}
	logx.AddGlobalFields(rs...)
}

func Alert(_ context.Context, v string) {
	logx.Alert(v)
}

func Close() error {
	return logx.Close()
}

func (l *DLogger) Debug(v ...interface{}) {
	l.print(l.logger.Debugw, v...)
}

func (l *DLogger) Debugf(format string, v ...interface{}) {
	l.printf(l.logger.Debugw, format, v...)
}

func (l *DLogger) Debugv(v interface{}) {
	l.printv(l.logger.Debugw, v)
}

func (l *DLogger) Debugw(msg string, fields ...LogField) {
	l.printw(l.logger.Debugw, msg, fields...)
}

func (l *DLogger) Error(v ...interface{}) {
	l.print(l.logger.Errorw, v...)
}

func (l *DLogger) Errorf(format string, v ...interface{}) {
	l.printf(l.logger.Errorw, format, v...)
}

func (l *DLogger) Errorv(v interface{}) {
	l.printv(l.logger.Errorw, v)
}

func (l *DLogger) Errorw(msg string, fields ...LogField) {
	l.printw(l.logger.Errorw, msg, fields...)
}

func (l *DLogger) Info(v ...interface{}) {
	l.print(l.logger.Infow, v...)
}

func (l *DLogger) Infof(format string, v ...interface{}) {
	l.printf(l.logger.Infow, format, v...)
}

func (l *DLogger) Infov(v interface{}) {
	l.printv(l.logger.Infow, v)
}

func (l *DLogger) Infow(msg string, fields ...LogField) {
	l.printw(l.logger.Infow, msg, fields...)
}

func (l *DLogger) Slow(v ...interface{}) {
	l.print(l.logger.Sloww, v...)
}

func (l *DLogger) Slowf(format string, v ...interface{}) {
	l.printf(l.logger.Sloww, format, v...)
}

func (l *DLogger) Slowv(v interface{}) {
	l.printv(l.Sloww, v)
}

func (l *DLogger) Sloww(msg string, fields ...LogField) {
	l.printw(l.logger.Sloww, msg, fields...)
}

func (l *DLogger) WithCallerSkip(skip int) *DLogger {
	l.logger.WithCallerSkip(skip)
	return l
}

func (l *DLogger) WithContext(ctx context.Context) *DLogger {
	l.logger.WithContext(ctx)
	return l
}

func (l *DLogger) WithDuration(d time.Duration) *DLogger {
	l.logger.WithDuration(d)
	return l
}

func (l *DLogger) WithFields(fields map[string]interface{}) *DLogger {
	clone := l.clone()
	for k, v := range fields {
		clone.fields = append(clone.fields, logx.Field(k, v))
	}
	return clone
}

func (l *DLogger) WithField(k string, v interface{}) *DLogger {
	clone := l.clone()

	clone.fields = append(clone.fields, logx.Field(k, v))

	return clone
}

func (l *DLogger) print(output func(msg string, fields ...logx.LogField), args ...interface{}) {
	output(fmt.Sprint(args...), l.fields...)
}

func (l *DLogger) printf(output func(msg string, fields ...logx.LogField), format string, args ...interface{}) {
	output(fmt.Sprintf(format, args...), l.fields...)
}

func (l *DLogger) printw(output func(msg string, fields ...logx.LogField), msg string, fields ...logx.LogField) {
	output(msg, append(l.fields, fields...)...)
}

func (l *DLogger) printv(output func(msg string, fields ...logx.LogField), msg interface{}) {
	output(fmt.Sprint(msg), l.fields...)
}

func (l *DLogger) clone() *DLogger {
	var fields []logx.LogField
	fields = append(fields, l.fields...)

	return &DLogger{
		logger: l.logger,
		fields: fields,
	}
}
