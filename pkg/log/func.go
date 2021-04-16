package log

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Errorw(format string, keysAndValues ...interface{}) {
	l.Errorw(format, keysAndValues...)
}

func Error(message string) {
	l.Error(message)
}

func Info(message string) {
	l.Info(message)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Infow(message string, keyAndValues ...interface{}) {
	l.Infow(message, keyAndValues...)
}
