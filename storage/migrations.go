package storage

import (
	chg "github.com/gophers-latam/challenges/http"
	"go.uber.org/zap"
)

func Migrate() {
	db := Get()

	err := db.AutoMigrate(chg.Command{}, chg.Challenge{}, chg.Fact{}, chg.Event{}, chg.Waifu{})
	if err != nil {
		zap.S().Fatal("cannot do migration, %s", err.Error())
	}
}
