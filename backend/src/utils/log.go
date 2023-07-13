package utils

import (
	"go.uber.org/zap"
)

func GetLog() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
