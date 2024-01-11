package nebula

import (
	"bytes"
	"io"
	"net/http"
)

type HttpLogger struct {
	*StdLogger
	host         string
	endpoint     string
	method       string
	headers      http.Header
	requestBody  []byte
	responseBody []byte
}

func (l *HttpLogger) SetRequest(req *http.Request) error {
	l.host = req.Host
	l.endpoint = req.URL.Path
	l.method = req.Method
	l.headers = req.Header
	l.requestBody = []byte("")
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		req.Body = io.NopCloser(bytes.NewReader(b))
		l.requestBody = b
	}

	return nil
}

func (l *HttpLogger) SetResponse(resp *http.Response) error {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body = io.NopCloser(bytes.NewReader(b))
	l.responseBody = b

	return nil
}

func (l *HttpLogger) Info() error {
	return push(
		l.nebulaHost,
		l.client,
		l.serviceName,
		httpLog,
		&log{
			level:   infoLogLevel,
			message: "",
			data:    l.constructHttpInfo(),
		},
	)
}

func (l *HttpLogger) Warning() error {
	return push(
		l.nebulaHost,
		l.client,
		l.serviceName,
		httpLog,
		&log{
			level:   warningLogLevel,
			message: "",
			data:    l.constructHttpInfo(),
		},
	)
}

func (l *HttpLogger) Error(err error) error {
	return push(
		l.nebulaHost,
		l.client,
		l.serviceName,
		httpLog,
		&log{
			level:   errorLogLevel,
			message: "",
			err:     err,
			data:    l.constructHttpInfo(),
		},
	)
}

func (l *HttpLogger) constructHttpInfo() map[string]any {
	return map[string]any{
		"host":         l.host,
		"endpoint":     l.endpoint,
		"method":       l.method,
		"headers":      l.headers,
		"requestBody":  string(l.requestBody),
		"responseBody": string(l.responseBody),
	}
}
