package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Connection struct{
	Conn	*gorm.DB
}


func Connect() *gorm.DB{
	dsn := "postgres://dev_user:saeedsaeed@postgres:5432/uncodev_db?sslmode=disable"
	db , err := gorm.Open(postgres.Open(dsn),&gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err!=nil {
		log.Fatal(err)
	}
	return db
}
