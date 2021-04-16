package noop

import (
	"github.com/nghiant3223/tikihackathon/pkg/log/logger"
)

var _ logger.Logger = (*noopLogger)(nil)

type noopLogger struct{}

func New() logger.Logger {
	return &noopLogger{}
}

func (n *noopLogger) Error(message string) {}

func (n *noopLogger) Info(message string) {}

func (n *noopLogger) Infof(format string, args ...interface{}) {}

func (n *noopLogger) Infow(message string, keyAndValues ...interface{}) {}

func (n *noopLogger) Panicf(format string, args ...interface{}) {}

func (n *noopLogger) Errorf(format string, args ...interface{}) {}

func (n *noopLogger) Errorw(message string, keyAndValues ...interface{}) {}
