package routes

import (
	"net/http"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/models"
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/utils"
	"github.com/gin-gonic/gin"
)

//signup parses the data it recieves and sends it to the 'Save' method.
func signup(con *gin.Context) {
	var user models.User

	err := con.ShouldBindJSON(&user)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	con.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

//login parses the data it recieves and sends it to the 'ValidateUser' method.
func login(con *gin.Context) {
	var user models.User

	err := con.ShouldBindJSON(&user)
	if err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		con.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateAuthToken(user.Email, user.ID)
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	con.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}