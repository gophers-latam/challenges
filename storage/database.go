package storage

import (
	"fmt"
	"os"
	"time"

	"github.com/gophers-latam/challenges/global"
	"gorm.io/driver/mysql"
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

// switch between local sqlite and mysql remote
func get() *gorm.DB {
	if env := os.Getenv("DBHOST"); env == "" {
		return getLocalDB()
	}
	return getRemoteDB()
}

func getRemoteDB() *gorm.DB {
	c := global.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DbUser, c.DbPass, c.DbHost, c.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbConfig, _ := db.DB()
	dbConfig.SetConnMaxIdleTime(1 * time.Hour)

	return db
}

func getLocalDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gophers.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
