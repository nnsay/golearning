package main

import (
	"time"

	"go.uber.org/zap"
)

func logWithSugaredLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}

func logWithZapLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Structured
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func main() {
	logWithZapLogger()

	logWithSugaredLogger()
}
