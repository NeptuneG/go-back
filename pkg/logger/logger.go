package logger

import (
	"log"

	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	return logger
}
