package database

import (
	"tgame/database/tables"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDatabase() {
	dsn := "host=postgres user=tgame password=tgame dbname=tgame port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error: cannot connect to the database")
	}

	db.AutoMigrate(&tables.User{}, &tables.Social{})

	DB = db
}
