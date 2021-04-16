package logger

type Logger interface {
	Panicf(format string, args ...interface{})
	Error(message string)
	Errorf(format string, args ...interface{})
	Errorw(message string, keyAndValues ...interface{})
	Info(message string)
	Infof(format string, args ...interface{})
	Infow(message string, keyAndValues ...interface{})
}
