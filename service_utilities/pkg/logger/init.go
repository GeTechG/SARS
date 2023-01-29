package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() error {
	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level, _ := zap.ParseAtomicLevel(os.Getenv("LOG_LEVEL"))

	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	sugarLogger = logger.Sugar()
	
	return nil
}

func GetLogger() *zap.Logger {
	return logger
}

func GetSugarLogger() *zap.SugaredLogger {
	return sugarLogger
}

func Flush() {
	_ = logger.Sync()
}
