package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var (
		checkProvinces bool
		checkDistricts bool
		checkUser      bool
	)

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&User{})

	checkProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Provinces Table")
	}

	checkDistricts = db.Migrator().HasTable(&Districts{})
	if !checkDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("Create Districts Table")
	}

	checkUser = db.Migrator().HasTable(&User{})
	if !checkUser {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Create User Table")
	}
}
