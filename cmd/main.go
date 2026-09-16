package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ryansissom/go-rest-api/internal/logger"
	"github.com/ryansissom/go-rest-api/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	r := router.New()
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(r))

	log.Info("server starting on port 8080")

	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
