package main

import (
	"mageBATestCase/docs"
	"mageBATestCase/routes"
	"mageBATestCase/util"
)

// @title  Mage BA Test Case API
// @description This is a test case for Mage BA position.Created by Ozgur Gurcan. phone number: +90 539 946 17 08
// @version  1.0

// @license MIT
// @contact.name Ozgur Gurcan
// @contact.email zgurgurcan1996@gmail.com

// @host      localhost:5050
// @BasePath  /api/v1

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5050"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "Mage BA Test Case API"
	docs.SwaggerInfo.Description = "This is a test case for Mage BA position."

	router := routes.InitRoute()

	port := util.GetEnvVariable("SERVER_PORT")
	router.Run(port)
}
