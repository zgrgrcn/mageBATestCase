package main

import (
	"mageBATestCase/routes"
	"mageBATestCase/util"
)

func main() {
	router := routes.InitRoute()
	port := util.GetEnvVariable("SERVER_PORT")
	router.Run(port)
}
