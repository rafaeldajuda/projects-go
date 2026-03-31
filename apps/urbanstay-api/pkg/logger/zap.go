package logger

import "go.uber.org/zap"

var lg *zap.Logger

func StartZapLog() {
	lg, _ = zap.NewProduction()
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
