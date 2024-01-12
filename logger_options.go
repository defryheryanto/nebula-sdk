package nebula

import "io"

type LoggerOption func(l *Logger)

func LoggerHostOption(host string) LoggerOption {
	return func(l *Logger) {
		l.host = host
	}
}

func LoggerWriterOption(w io.Writer) LoggerOption {
	return func(l *Logger) {
		l.w = w
	}
}
