package entity

import "gorm.io/gorm"

type User struct {
	Uuid     string `gorm:"column:uuid;primaryKey"`
	Email    string `gorm:"column:email;index;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Verified bool   `gorm:"column:verified;not null;default:false"`
}

func UserTable(user User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("user")
	}
}
