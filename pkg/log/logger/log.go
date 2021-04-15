package log

import (
	"github.com/nghiant3223/tikihackathon/pkg/log/zap"
	"sync"
)

var (
	l Logger
	o sync.Once
)

func Init() {
	l = zap.New()
}

type Logger interface {
	Panicf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Errorw(message string, keyAndValues ...interface{})
}
