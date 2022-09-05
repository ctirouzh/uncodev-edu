package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Connection struct{
	Conn	*gorm.DB
}



func Connect(){

	dsn := "postgres://saeed:123456@127.0.0.1:5432/crypto?TimeZone=Asia/Tehran"
	
	db , err := gorm.Open(postgres.Open(dsn),&gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err!=nil {
		log.Fatal(err)
	}

}