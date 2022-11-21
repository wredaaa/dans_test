package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

/*
	Do this function to connect database.

*
*/
func GetDatabase() *gorm.DB {
	connection, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dans?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful.")
	return connection
}

/*
	Do this function to create user table in userdb.

*
*/
func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(User{})
}

/*
	Do this function to close database connection.

*
*/
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
