package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// Db is a function to connect to database
func Db() *gorm.DB {
	// change this database configuration according to your own
	db, err := gorm.Open(mysql.Open("username:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}