package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Get() *gorm.DB {
	if DB == nil {
		DB = get()
	}
	return DB
}

func get() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gophers.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
