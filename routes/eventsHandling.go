package routes

import (
	"net/http"
	"strconv"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/models"
	"github.com/gin-gonic/gin"
)

//getEvents returns all events in the database.
func getEvents(con *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}
	con.JSON(http.StatusOK, events)
}

//getEvent returns the event specified by the 'id' parameter.
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

//createEvent recieves the event data and sends it to the 'Save' method.
func createEvent(con *gin.Context) {
	var event models.Event

	err := con.ShouldBindJSON(&event)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.UserID = con.GetInt64("userId")

	err = event.Save()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save evemt"})
		return
	}

	con.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

//updateEvent recieves the updated event data and the 'id' parameter and sends them to the 'UpdateEvent' method.
func updateEvent(con *gin.Context) {
	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	userId := con.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not locate event"})
		return
	}

	if event.UserID != userId {
		con.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
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

//deleteEvent sends the event specified by the 'id' paramater to the DeleteEvent method.
func deleteEvent(con *gin.Context) {
	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	userId := con.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not locate event"})
		return
	}

	if event.UserID != userId {
		con.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}

	con.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
