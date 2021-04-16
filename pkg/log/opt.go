package log

import (
	"github.com/nghiant3223/tikihackathon/pkg/log/logger"
)

type opt struct {
	mode       logger.Mode
	loggerType logger.Type
}

type optFn func(*opt)

func WithLoggerType(t string) optFn {
	return func(o *opt) {
		o.loggerType = logger.Atot(t)
	}
}

func WithMode(mode string) optFn {
	return func(o *opt) {
		o.mode = logger.Atom(mode)
	}
}
