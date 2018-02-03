package main

import (
	// go orm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-academy/user_auth/model"
)

// returns db connection and error
func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"user=carmento password=carmento dbname=go_user_auth sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	// migrates creating table for all your models
	db.AutoMigrate(&model.User{}, &model.Message{})

	return db, nil
}
