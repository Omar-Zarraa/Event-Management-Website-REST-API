package routes

import (
	"net/http"

	"github.com/Omar-Zarraa/REST-API/models"
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
