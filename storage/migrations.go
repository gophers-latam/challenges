package storage

import (
	"github.com/tomiok/challenge-svc/challenges"
	"go.uber.org/zap"
)

func Migrate() {
	db := Get()

	err := db.AutoMigrate(challenges.Challenge{})

	if err != nil {
		zap.S().Fatal("cannot do migration, %s", err.Error())
	}
}
