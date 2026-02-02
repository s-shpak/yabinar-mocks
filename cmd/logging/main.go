package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	var example string
	flag.StringVar(&example, "example", "", "Logging example to run")

	flag.Parse()

	if len(example) == 0 {
		log.Fatal("example must be specified")
	}

	switch example {
	case "simple":
		simple()
	case "format-attrs":
		formatAttrs()
	case "levels":
		levels()
	case "nested":
		nested()
	case "ctx":
		ctx()
	default:
		log.Fatalf("unknown example %q", example)
	}
}

func simple() {
	slog.Info("Hello!", "receiver", "world")

	localLogger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	localLogger.Info("This is printed to stderr", "logger", "local")

	localLogger.Info(
		"This is printed to stderr",
		slog.Int("answer", 42),
	)
}

func formatAttrs() {
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()

				a.Value = slog.StringValue(t.Format(time.RFC3339))
				return a
			}

			if a.Key == "logger" {
				a.Value = slog.StringValue("the value has been modified")
				return a
			}

			return a
		},
	}
	localLogger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	localLogger.Info("Hello, world!", "logger", "local")
}

func levels() {
	loggerLevel := new(slog.LevelVar)
	loggerLevel.Set(slog.LevelInfo)
	opts := &slog.HandlerOptions{
		Level: loggerLevel,
	}

	localLogger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	localLogger.Debug("Debug msg")
	localLogger.Info("Info msg")

	loggerLevel.Set(slog.LevelDebug)

	localLogger.Debug("Second debug msg")
	localLogger.Info("Second info msg")
}

func nested() {
	localLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	localLogger.Info(
		"This is a message with a nested object",
		slog.Group(
			"request",
			"endpoint", "/api/v1/log",
			"method", http.MethodGet,
		),
	)
}

func ctx() {
	localLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ctx := context.WithValue(
		context.Background(),
		"spanID", "12345",
	)

	localLogger.InfoContext(ctx, "Msg without context", "component", "handler")

	localContextLogger := slog.New(NewContextHandler(localLogger.Handler()))

	localContextLogger.InfoContext(ctx, "Msg with context", "component", "handler")
}

type ContextHandler struct {
	slog.Handler
}

func NewContextHandler(h slog.Handler) *ContextHandler {
	return &ContextHandler{Handler: h}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if spanID, ok := ctx.Value("spanID").(string); ok {
		r.AddAttrs(slog.String("spanID", spanID))
	}
	return h.Handler.Handle(ctx, r)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithGroup(name)}
}
