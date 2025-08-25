//Package routes handles the requests coming to the API.
package routes

import (
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/middlewares"
	"github.com/gin-gonic/gin"
)

//RegisterRoutes specifies which function to call on for each request/route pair.
//(GET/'/events'): returns all events, (GET/'/event/:id'): returns the specified event, (POST/'/events'): create an event,
//(PUT/'/events/:id'): update specified event, (DELETE/'/events/:id'): delete specified event, (POST/'/events/:id/register'): register user for specified event,
//(DELETE/'/events/:id/register'): cancel the registration on specified event, (POST'/signup'): create a user account, (POST/'/login'): log into user account.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
