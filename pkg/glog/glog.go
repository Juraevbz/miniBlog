package glog

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

const logDir = "logs"

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId := getSpanIdFromContext(ctx)
	e.Str("span-id", spanId)
}

type TimestampHook struct{}

func (t TimestampHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Time("datetime", time.Now().UTC())
}

func getSpanIdFromContext(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	return span.SpanContext().SpanID().String()
}

func NewLogger() zerolog.Logger {
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	errorLog, err := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open error log file: " + err.Error())
	}

	infoLog, err := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open info log file: " + err.Error())
	}

	warnLog, err := os.OpenFile(filepath.Join(logDir, "warn.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open warn log file: " + err.Error())
	}

	debugLog, err := os.OpenFile(filepath.Join(logDir, "debug.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open debug log file: " + err.Error())
	}

	multiError := zerolog.MultiLevelWriter(errorLog)
	multiInfo := zerolog.MultiLevelWriter(infoLog)
	multiWarn := zerolog.MultiLevelWriter(warnLog)
	multiDebug := zerolog.MultiLevelWriter(debugLog)

	logger := zerolog.New(zerolog.MultiLevelWriter(multiError, multiInfo, multiWarn, multiDebug)).Hook(TimestampHook{}).With().Logger()

	return logger
}
