package nebula

import "io"

type LoggerOption func(l *Logger)

func LoggerHostOption(host string) LoggerOption {
	return func(l *Logger) {
		l.host = host
	}
}

func LoggerStdWriterOption(w io.Writer) LoggerOption {
	return func(l *Logger) {
		l.stdWriter = w
	}
}

func LoggerHttpWriterOption(w io.Writer) LoggerOption {
	return func(l *Logger) {
		l.httpWriter = w
	}
}
