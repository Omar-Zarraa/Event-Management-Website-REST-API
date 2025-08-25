//Package main is the package that launches the API.
package main

import (
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/db"
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/routes"
	"github.com/gin-gonic/gin"
)

//main initializes the database and server and sets it to listen on port 8080, and calls on the routes package.
func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
