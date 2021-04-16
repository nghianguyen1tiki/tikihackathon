package zap

import (
	"go.uber.org/zap"

	"github.com/nghiant3223/tikihackathon/pkg/log/logger"
)

var _ logger.Logger = (*zapLogger)(nil)

type zapLogger struct {
	lg        *zap.Logger
	sugaredLg *zap.SugaredLogger
}

func New(optFns ...optFn) (*zapLogger, error) {
	zapLg := &zapLogger{}
	o := &opt{}
	for _, optFn := range optFns {
		optFn(o)
	}
	err := zapLg.initialize(o)
	if err != nil {
		return nil, err
	}
	return zapLg, nil
}

func (l *zapLogger) initialize(o *opt) error {
	var lg *zap.Logger
	var err error

	switch o.mode {
	case logger.Production:
		lg, err = zap.NewProduction()
	case logger.Development:
		lg, err = zap.NewDevelopment()
	}
	if err != nil {
		return err
	}
	defer lg.Sync()

	l.lg = lg
	l.sugaredLg = lg.Sugar()

	return nil
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLg.Panicf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLg.Errorf(format, args...)
}

func (l *zapLogger) Errorw(message string, keyAndValues ...interface{}) {
	l.sugaredLg.Errorw(message, keyAndValues...)
}

func (l *zapLogger) Error(message string) {
	l.lg.Error(message)
}

func (l *zapLogger) Info(message string) {
	l.lg.Info(message)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLg.Infof(format, args...)
}

func (l *zapLogger) Infow(message string, keyAndValues ...interface{}) {
	l.sugaredLg.Infow(message, keyAndValues...)
}
