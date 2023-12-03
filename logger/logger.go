package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return Logger
}

func GetLoggerSugar() *zap.SugaredLogger {
	return Logger.Sugar()
}
