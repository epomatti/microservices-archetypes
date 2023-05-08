package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	var err error
	db, err = gorm.Open(sqlite.Open(dns()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
}

const (
	File   string = "test.db"
	Memory string = ":memory:"
)

func dns() string {
	switch os.Getenv("SQLITE_DSN") {
	case "file":
		return File
	case "memory":
		return Memory
	default:
		return File
	}
}
