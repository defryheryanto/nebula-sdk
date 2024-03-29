package nebula

import (
	"io"
	"net/http"
)

type StdLogger struct {
	nebulaHost  string
	serviceName string
	client      *http.Client
	w           io.Writer
}

func (l *StdLogger) Info(msg string, data map[string]any) error {
	return push(
		l.w,
		l.nebulaHost,
		l.client,
		l.serviceName,
		stdLog,
		&log{
			level:   infoLogLevel,
			message: msg,
			data:    data,
		},
	)
}

func (l *StdLogger) Warning(msg string, data map[string]any) error {
	return push(
		l.w,
		l.nebulaHost,
		l.client,
		l.serviceName,
		stdLog,
		&log{
			level:   warningLogLevel,
			message: msg,
			data:    data,
		},
	)
}

func (l *StdLogger) Error(msg string, err error, data map[string]any) error {
	return push(
		l.w,
		l.nebulaHost,
		l.client,
		l.serviceName,
		stdLog,
		&log{
			level:   errorLogLevel,
			message: msg,
			err:     err,
			data:    data,
		},
	)
}
