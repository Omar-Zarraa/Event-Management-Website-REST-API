package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080") //localhost:8080
}

func getEvents(con *gin.Context) {
	con.JSON(http.StatusOK, gin.H{"message": "Done!"})
}
