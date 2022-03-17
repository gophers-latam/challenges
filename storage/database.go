package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

var DB *gorm.DB

func Get() *gorm.DB {
	if DB == nil {
		DB = get()
	}
	return DB
}

func get() *gorm.DB {
	if env := os.Getenv("ENV"); env == "" || env == "local" {
		return getSqlite()
	}
	return getClearDB()
}

func getClearDB() *gorm.DB {
	dsn := os.Getenv("CLEARDB_DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbConfig, _ := db.DB()
	dbConfig.SetMaxIdleConns(0)
	dbConfig.SetMaxIdleConns(10)
	dbConfig.SetConnMaxIdleTime(1 * time.Hour)

	return db
}

func getSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gophers.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
