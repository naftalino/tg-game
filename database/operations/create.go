package operations

import (
	"tgame/database/tables"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *tables.User) error {
	return db.Create(user).Error
}
