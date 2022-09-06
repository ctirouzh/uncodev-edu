package config

import (
	"log"
	"github.com/joho/godotenv"
)



func SetConfig(){

	err := godotenv.Load("./config/.env")
	if err!=nil{
		log.Fatalf("Some error occured. Err %s",err)
	}

}