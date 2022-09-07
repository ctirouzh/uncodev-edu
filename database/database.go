package database

import (
	"fmt"
	"log"
	"github.com/uncodev/educational-website-with-golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)


type Connection struct{
	Conn	*gorm.DB
}


var GetConnection Connection

var doOnce sync.Once

func Connect() *gorm.DB {
	doOnce.Do(func() {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v timezone=%v",
			config.AppConfig.DbConfig.Host,
			config.AppConfig.DbConfig.User,
			config.AppConfig.DbConfig.Password,
			config.AppConfig.DbConfig.Name,
			config.AppConfig.DbConfig.TimeZone,
		)
		db , err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err!=nil {
			log.Fatal(err)
		}
		GetConnection.Conn = db
	})
	return GetConnection.Conn
}

