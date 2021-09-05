package logs

import "go.uber.org/zap"

func InitLogs() {
	logger, err := zap.NewProduction()
	_ = logger.Sync()
	if err != nil {
		panic(err.Error())
	}

	zap.ReplaceGlobals(logger)
}
