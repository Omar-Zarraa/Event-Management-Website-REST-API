package routes

import (
	"net/http"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/models"
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/utils"
	"github.com/gin-gonic/gin"
)

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

func getUsers(con *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		con.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get users"})
		return
	}
	con.JSON(http.StatusOK, users)
}
