package routes

import (
	"net/http"
	"strconv"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/models"
	"github.com/gin-gonic/gin"
)

//registerForEvent gets the userId of the user and the id of the event from the id parameter and sends them to the 'Register' method.
func registerForEvent(con *gin.Context) {
	userId := con.GetInt64("userId")

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

	err = event.Register(userId)
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	con.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

//cancelRegistration gets the userId of the user and the id of the event from the id parameter and sends them to the 'CancelRegistration' method.
func cancelRegistration(con *gin.Context) {
	userId := con.GetInt64("userId")

	eventId, err := strconv.ParseInt(con.Param("id"), 10, 64)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	con.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}
