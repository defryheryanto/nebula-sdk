package nebula

type LoggerOption func(l *Logger)

func LoggerHostOption(host string) LoggerOption {
	return func(l *Logger) {
		l.host = host
	}
}
