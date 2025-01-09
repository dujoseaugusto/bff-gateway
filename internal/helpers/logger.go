package helpers

import (
	"context"
	"log"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

type Logger struct {
	*log.Logger
}

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{logger}
}

func (l *Logger) WithContext(ctx context.Context) *log.Logger {
	requestID, ok := ctx.Value(RequestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}
	return log.New(l.Writer(), l.Prefix()+" [requestID="+requestID+"] ", l.Flags())
}
