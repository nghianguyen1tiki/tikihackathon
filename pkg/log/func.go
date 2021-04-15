package log

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Errorw(format string, args ...interface{}) {
	l.Errorw(format, args...)
}
