package observability

import (
	"fmt"
	"time"
)

type Severity int

const (
	INFO Severity = iota
	WARNING
	ERROR
	FATAL
)

type ErrorHeader struct {
	occuredAt time.Time
	severity  Severity
}

func NewErrorHeader(severity Severity) ErrorHeader {
	return ErrorHeader{
		occuredAt: time.Now(),
		severity:  severity,
	}
}

func (h ErrorHeader) SeverityName() string {
	switch h.severity {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type Error struct {
	header     ErrorHeader
	message    string
	parameters []any
}

func newError(severity Severity, message string, parameters ...any) *Error {
	return &Error{
		header:     NewErrorHeader(severity),
		message:    message,
		parameters: parameters,
	}
}

func NewInformationalMessage(name, message string, parameters ...any) *Error {
	return newError(INFO, message, parameters...)
}

func NewWarning(name, message string, parameters ...any) *Error {
	return newError(WARNING, message, parameters...)
}

func NewError(severity Severity, message string, parameters ...any) *Error {
	return newError(ERROR, message, parameters...)
}

func NewFatalError(name, message string, parameters ...any) *Error {
	return newError(FATAL, message, parameters...)
}

func (e *Error) Errorf() error {
	return fmt.Errorf("%s : %s > %s", e.header.occuredAt.Format(time.RFC3339), e.header.SeverityName(),
		fmt.Sprintf("%s : %v >"+e.message, e.parameters...))
}
