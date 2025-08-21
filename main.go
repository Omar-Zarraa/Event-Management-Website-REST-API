package main

import (
	"github.com/Omar-Zarraa/REST-API/db"
	"github.com/Omar-Zarraa/REST-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
