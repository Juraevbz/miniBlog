package glog

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId := getSpanIdFromContext(ctx)
	e.Str("span-id", spanId)
}

func getSpanIdFromContext(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	return span.SpanContext().SpanID().String()
}

func NewTracingLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	logger = logger.Hook(TracingHook{})
	return logger
}

func NewLogger() zerolog.Logger {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(err)
	}

	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02") + ".log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	multi := zerolog.MultiLevelWriter(file, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger
}
