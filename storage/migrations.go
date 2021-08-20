package storage

import (
	"github.com/tomiok/challenge-svc/challenges"
	"log"
)

func Migrate() {
	db := Get()

	err := db.AutoMigrate(challenges.Challenge{})

	if err != nil {
		log.Print(err.Error())
		panic("cannot do migration :( ")
	}
}
