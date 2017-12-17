package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	url := "holy url"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url",
		"attempt",
		"backoff", time.Second)
	sugar.Infof("Failed to fetch URL: %s", url)
}
