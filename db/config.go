package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var gormConn *gorm.DB

// GetDatabaseConnection get connection database
func GetDatabaseConnection() *gorm.DB { // Check Connection Status
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	conn, err := gorm.Open("postgres", "host=localhost port=5432 user=test dbname=test password=test sslmode=disable")

	if err != nil {
		panic(err) // log error without close
	}
	conn.LogMode(true)
	gormConn = conn
	return gormConn
}
