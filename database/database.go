package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Connection struct{
	Conn	*gorm.DB
}


var Conn Connection

func Connect(){
	dsn := "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable"
	db , err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil {
		log.Fatal(err)
	}
	Conn.Conn = db	
}