package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "host=postgres user=postgres password=postgres dbname=go_crud_app port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
