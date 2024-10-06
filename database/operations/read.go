package operations

import (
	"tgame/database/tables"

	"gorm.io/gorm"
)

func FindUserByID(db *gorm.DB, ID uint64) (*tables.User, error) {
	var user tables.User

	err := db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func TotalUsers(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&tables.User{}).Count(&count).Error
	return count, err
}

func UserWithMoreDraws(db *gorm.DB) (*tables.User, error) {
	var user tables.User

	err := db.Order("total_draws desc").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func UserWithMoreCoins(db *gorm.DB) (*tables.User, error) {
	var user tables.User

	err := db.Order("coins desc").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
