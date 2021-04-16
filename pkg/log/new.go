package log

import (
	"sync"

	"github.com/nghiant3223/tikihackathon/pkg/log/logger"
	"github.com/nghiant3223/tikihackathon/pkg/log/noop"
	"github.com/nghiant3223/tikihackathon/pkg/log/zap"
)

var (
	l    logger.Logger
	once sync.Once
)

func init() {
	Init()
}

func Init(optFns ...optFn) {
	once.Do(func() {
		o := &opt{
			mode:       logger.Production,
			loggerType: logger.Zap,
		}
		for _, fn := range optFns {
			fn(o)
		}
		l = newLogger(o)
	})
}

func newLogger(o *opt) logger.Logger {
	switch o.loggerType {
	case logger.Zap:
		l, err := zap.New(zap.WithMode(o.mode))
		if err == nil {
			return l
		}
		fallthrough
	case logger.NoOp:
		fallthrough
	default:
		return noop.New()
	}
}
