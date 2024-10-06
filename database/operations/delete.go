package operations

import (
	"tgame/database/tables"

	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, user *tables.User) error {
	return db.Delete(user).Error
}
