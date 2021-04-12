package entity

import "gorm.io/gorm"

type RefreshToken struct {
	Uuid       string `gorm:"column:uuid;primaryKey"`
	UserUuid   string `gorm:"column:user_uuid;not null"`
	Expiration int64  `gorm:"column:expiration;not null"`
}

func RefreshTokenTable(refreshToken RefreshToken) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("refresh_token")
	}
}
