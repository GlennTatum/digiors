package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&Station{},
		&Registration{},
		&AnonUser{},
	)
}
