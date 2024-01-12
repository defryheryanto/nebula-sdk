package nebula

var l *Logger

func init() {
	l = NewLogger("unspecified")
}

func SetLogger(logger *Logger) {
	l = logger
}

func Std() *StdLogger {
	return l.Std()
}

func Http() *HttpLogger {
	return l.Http()
}
