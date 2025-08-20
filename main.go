package main

import (
	"net/http"

	"github.com/Omar-Zarraa/REST-API/db"
	"github.com/Omar-Zarraa/REST-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(con *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	con.JSON(http.StatusOK, events)
}

func createEvent(con *gin.Context) {
	var event models.Event

	err := con.ShouldBindJSON(&event)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.ID = 1

	err = event.Save()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	con.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
