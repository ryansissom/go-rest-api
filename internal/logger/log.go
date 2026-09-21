package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// CtxKey identifies the request logger stored in a context.
type CtxKey struct{}

// CtxWithLogger returns a context containing logger when logger is non-nil.
func CtxWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	if logger == nil {
		return ctx
	}

	if ctxLog, ok := ctx.Value(CtxKey{}).(*slog.Logger); ok && ctxLog == logger {
		return ctx
	}

	return context.WithValue(ctx, CtxKey{}, logger)
}

// FromContext returns the request logger or a default stdout logger.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(CtxKey{}).(*slog.Logger); ok {
		return logger
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}

// AddLoggerMid places the application logger into the request context.
func AddLoggerMid(logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggerCtx := CtxWithLogger(r.Context(), logger)
		r = r.Clone(loggerCtx)
		next.ServeHTTP(w, r)
	}
}

// LoggerMid logs the request path before passing control to the next handler.
func LoggerMid(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := FromContext(r.Context())
		l.Info("request", "path", r.URL.String())
		next.ServeHTTP(w, r)
	}
}
