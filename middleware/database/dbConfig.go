package database

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBConn Current Database connection
	DBConn *gorm.DB
	// GODB
	DBGo *sql.DB
	// Err Database connection error
	Err error
)

func PostgresConfig(username, password, host, databaseName, port, sslMode, timeZone string) {
	DBConn, Err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			host, username, password, databaseName, port, sslMode, timeZone)),
		&gorm.Config{})
}

func ConnectDatabase() {
	PostgresConfig(
		"postgres",    // Username
		"1Qaz2wsx",    // Password
		"localhost",   // Host
		"fds_api",     // DB Name
		"5432",        // Port
		"disable",     // SSL
		"Asia/Manila", //Timezone
	)
	err := DBConn.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}
}
