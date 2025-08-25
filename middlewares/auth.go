//Package middlewares contains the middleware used for password authentication.
package middlewares

import (
	"net/http"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/utils"
	"github.com/gin-gonic/gin"
)

//Authenticate gets the authentication token sent with the user and sends it to the 'VerifyAuthToken' function.
func Authenticate(con *gin.Context) {
	token := con.Request.Header.Get("Authorization")

	if token == "" {
		con.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyAuthToken(token)
	if err != nil {
		con.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	con.Set("userId", userId)
	con.Next()
}
