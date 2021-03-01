package connection

import (
	"fmt"
	"graphql-intro/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect to database
func Connect() {

	var userDB, pwDB, portDB, hostDB, nameDB, sslDB, timeZoneDB string
	userDB = helper.GetEnvVar("DB_USER")
	pwDB = helper.GetEnvVar("DB_PASSWORD")
	portDB = helper.GetEnvVar("DB_PORT")
	hostDB = helper.GetEnvVar("DB_HOST")
	nameDB = helper.GetEnvVar("DB_NAME")
	sslDB = helper.GetEnvVar("DB_SSL")
	timeZoneDB = helper.GetEnvVar("DB_TIMEZONE")

	conn := " host=" + hostDB +
		" user=" + userDB +
		" password=" + pwDB +
		" dbname=" + nameDB +
		" port=" + portDB +
		" sslmode=" + sslDB +
		" TimeZone=" + timeZoneDB

	db, errConn := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to the database")
	} else {
		fmt.Println("successful connection")
	}
	dbConn = db
}

//GetConnection from DB
func GetConnection() *gorm.DB {
	return dbConn
}
