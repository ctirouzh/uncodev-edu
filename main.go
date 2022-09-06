package main

import (
	"github.com/uncodev/educational-website-with-golang/config"
	"github.com/uncodev/educational-website-with-golang/database"
	"github.com/uncodev/educational-website-with-golang/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){

	// set config
	config.SetConfig()

	// set database
	database.Connect()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))  
	
	e.Static("/", "public")

	// set routes
	routes.SetRoutes(e)

	e.Logger.Fatal(e.Start(":9090"))

}