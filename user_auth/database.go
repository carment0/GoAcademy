package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"user=carmento password=carmento dbname=go_user_auth sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	// creating table for all your models
	//db.AutoMigrate()

	return db, nil
}
