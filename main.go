package main

import (
	"log"
	"mageBATestCase/docs"
	"mageBATestCase/routes"
)

// @title  Mage BA Test Case API
// @description This is a test case for Mage BA position.Created by Ozgur Gurcan. phone number: +90 539 946 17 08
// @version  1.0

// @license MIT
// @contact.name Ozgur Gurcan
// @contact.email zgurgurcan1996@gmail.com

// @host      https://shielded-springs-54397.herokuapp.com/
// @BasePath  /api/v1

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://shielded-springs-54397.herokuapp.com/"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "Mage BA Test Case API"
	docs.SwaggerInfo.Description = "This is a test case for Mage BA position."

	router := routes.InitRoute()

	//port := util.GetEnvVariable("PORT")
	err := router.Run()
	if err != nil {
		log.Fatal("Error while running the server: ", err)
	}
}
