package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	account "with.orm/services/account/domain"
)

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&account.Account{})

	return db
}
