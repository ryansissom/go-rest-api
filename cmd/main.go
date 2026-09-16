package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ryansissom/go-rest-api/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	logger.Info("server starting on port 8080")
	if err := http.ListenAndServe(":8080", router.New()); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
