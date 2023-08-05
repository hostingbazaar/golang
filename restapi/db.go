package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// root@localhost:3306
var Database *gorm.DB
var urlDSN = "root:admin@tcp(localhost:3306)/restapi"
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
