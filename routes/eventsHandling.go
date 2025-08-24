package routes

import (
	"net/http"
	"strconv"

	"github.com/Omar-Zarraa/REST-API/models"
	"github.com/gin-gonic/gin"
)

func getEvents(con *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}
	con.JSON(http.StatusOK, events)
}

func getEvent(con *gin.Context) {
	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	con.JSON(http.StatusOK, event)
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
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save evemt"})
		return
	}

	con.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(con *gin.Context) {
	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not locate event"})
		return
	}

	var updatedEvent models.Event

	err = con.ShouldBindJSON(&updatedEvent)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	con.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(con *gin.Context) {
	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not locate event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}

	con.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
