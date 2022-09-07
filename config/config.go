package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct{
	DbConfig	databse
}

type databse struct{
	Port		string
	Host		string
	User		string
	Password	string
	Name		string
	TimeZone	string
}

var AppConfig config

func SetConfig(){
	err := godotenv.Load("./config/.env")
	if err!=nil{
		log.Fatalf("Some error occured. Err %s",err)
	}
	AppConfig.DbConfig.Host = os.Getenv("POSTGRES_HOST")
	AppConfig.DbConfig.Port = os.Getenv("POSTGRES_PORT")
	AppConfig.DbConfig.User = os.Getenv("POSTGRES_USER")
	AppConfig.DbConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	AppConfig.DbConfig.Name = os.Getenv("POSTGRES_NAME")
	AppConfig.DbConfig.TimeZone = os.Getenv("POSTGRES_TIMEZONE")
}