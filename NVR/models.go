package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Camera struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string
	URL           string
	RetentionDays int
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("nvr.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Camera{})
}
