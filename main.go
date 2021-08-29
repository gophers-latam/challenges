package main

import (
	"github.com/tomiok/challenge-svc/logs"
	"github.com/tomiok/challenge-svc/storage"
	"go.uber.org/zap"
	"os"
)

func main() {
	logs.InitLogs()
	storage.Migrate()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	zap.L().Fatal("cannot init server", zap.Error(start(port)))
}
