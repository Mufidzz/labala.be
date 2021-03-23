package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ofcode.dev/labala-backend/structs"
)

const username = "root"
const password = ""

//const username = "atlok"
//const password = "f67f589174f0d2866ba8f3b40286d72e"
const dbName = "dev.ofcode.labala"

func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(0)

	err = db.AutoMigrate(structs.User{}, structs.Product{}, structs.ProductExchange{})

	if err != nil {
		panic(err)
	}

	return db
}
