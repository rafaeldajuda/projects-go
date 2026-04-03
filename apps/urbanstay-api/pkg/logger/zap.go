package logger

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

var lg *zap.Logger

func StartZapLog() {
	logLevel := os.Getenv("LOG_LEVEL")
	fmt.Println("SET LOG_LEVEL:", logLevel)
	config := zap.NewProductionConfig()

	switch logLevel {
	case "DEBUG":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "INFO":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "ERROR":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	lg, _ = config.Build()
}

func Sync() error {
	return lg.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	lg.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	lg.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	lg.Error(msg, fields...)
}

func JsonToString(obj interface{}) string {
	s, _ := json.Marshal(obj)
	return string(s)
}
