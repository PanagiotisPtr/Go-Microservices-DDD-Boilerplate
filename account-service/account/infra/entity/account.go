package entity

import "gorm.io/gorm"

type Account struct {
	Uuid     string `gorm:"column:uuid;primaryKey"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func AccountTable(account Account) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("account")
	}
}
