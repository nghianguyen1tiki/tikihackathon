package zap

import (
	"github.com/nghiant3223/tikihackathon/pkg/logger"
)

var _ logger.Logger = (*zapLogger)(nil)

type zapLogger struct{}

func NewZapLogger() *zapLogger {
	return &zapLogger{}
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {}

func (l *zapLogger) Errorf(format string, args ...interface{}) {}

func (l *zapLogger) Errorw(message string, keyAndValues ...interface{}) {}
