package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// root@localhost:3306
var Database *gorm.DB
var urlDSN = "sql6638305:xijJ4pwZZS@tcp(sql6.freesqldatabase.com:3306)/sql6638305"
var err error

func Datamigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		// return nil,err
		fmt.Println(err.Error())
		panic("Connection failed :(")
	}
	Database.AutoMigrate(&Employee{})
}
