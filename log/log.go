package log

import (
	"go.uber.org/zap"
)

func Logger() *zap.SugaredLogger {
	// logger, _ := zap.NewDevelopment()
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar
}
