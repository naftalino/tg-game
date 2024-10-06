package tables

type User struct {
	ID          *uint64  `gorm:"unique"`
	Draws       *int     `gorm:"default:10"`
	TotalDraws  *uint    `gorm:"default:0"`
	Coins       *int     `gorm:"default:0"`
	Bio         *string  `gorm:"default:'Eu ainda não mudei isso.'"`
	SocialMidia []Social `gorm:"foreignKey:UserId;references:ID"`
}

type Social struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	UserID   uint64
	Platform string
	Username string
}
