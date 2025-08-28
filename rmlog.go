package rmlog

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/aeternitas-infinita/rmlog/pkg/core"
	"github.com/aeternitas-infinita/rmlog/pkg/handler"
)

var Log = slog.New(handler.NewCustomHandler(
	os.Stdout,
	core.GetLvlFromEnv("log_level"),
	true,
	false,
))

type LoggerConfig struct {
	Level         slog.Level
	SentryEnabled bool
	TraceIDKey    string
	AddSource     bool
}

func CreateLogger(config LoggerConfig) *slog.Logger {
	handler := handler.NewCustomHandler(os.Stdout, config.Level, config.AddSource, config.SentryEnabled)
	return slog.New(handler)
}

func TraceIDToFHCtx(ctx *fasthttp.RequestCtx) {
	core.TraceIDToFHCtx(ctx)
}

func CtxWithTraceID(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return core.CtxWithTraceID(parent, timeout)
}

func GetTraceID(ctx any) string {
	return core.GetTraceID(ctx)
}

func ErrAttr(err error) slog.Attr {
	return core.ErrAttr(err)
}

func GetLvlFromStr(s string) slog.Level {
	return core.GetLvlFromStr(s)
}

func UpdateTraceIDKey(s string) {
	core.TraceIDKey = s
}

func GetBoolFromStr(s string) bool {
	return core.GetBoolFromStr(s)

}
