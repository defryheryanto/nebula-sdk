package nebula

var l *Logger

func init() {
	l = NewLogger("unspecified")
}

func SetLogger(logger *Logger) {
	l = logger
}

func StdLog() *StdLogger {
	return l.Std()
}

func HttpLog() *HttpLogger {
	return l.Http()
}
