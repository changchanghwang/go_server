package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	product "with.orm/services/products/domain"
)

func Init() *gorm.DB {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
			LogLevel:                  logger.Info, // Log level
		},
	)
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=UTC"), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&product.Product{})

	return db
}
