package storage

import (
	"github.com/gophers-latam/challenges/challenges"
	"go.uber.org/zap"
)

func Migrate() {
	db := Get()

	err := db.AutoMigrate(challenges.Challenge{})
	if err != nil {
		zap.S().Fatal("cannot do migration, %s", err.Error())
	}
}
