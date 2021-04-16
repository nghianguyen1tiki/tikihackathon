package zap

import "github.com/nghiant3223/tikihackathon/pkg/log/logger"

type opt struct {
	mode logger.Mode
}

type optFn func(*opt)

func WithMode(mode logger.Mode) optFn {
	return func(o *opt) {
		o.mode = mode
	}
}
