package log

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugz(msg string, fields ...zap.Field)

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoz(msg string, fields ...zap.Field)

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnz(msg string, fields ...zap.Field)

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorz(msg string, fields ...zap.Field)

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalz(msg string, fields ...zap.Field)

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger

	SetLogLevel(level Level) error
}

func New(Type LoggerType, opts ...Option) (Logger, error) {
	switch Type {
	case ZapLogger:
		return newZapLogger(opts...)
	default:
		return nil, fmt.Errorf("Invaild LoggerType:%v ", Type)
	}
}
