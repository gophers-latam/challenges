package logs

import "go.uber.org/zap"

func InitLogs() {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err.Error())
	}

	zap.ReplaceGlobals(logger)
}