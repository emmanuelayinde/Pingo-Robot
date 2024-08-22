package main

import (
	"github.com/emmanuelayinde/pingo-robot/config"
	"github.com/emmanuelayinde/pingo-robot/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080") // Default listens on port 8080
}
