package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkProvinces bool

	db.Migrator().DropTable(&Provinces{})

	checkProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Provinces Table")
	}

}
