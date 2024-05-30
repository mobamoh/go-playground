package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", GetEvents)
	r.GET("/events/:id", GetEvent)

	rgAuth := r.Group("/")
	rgAuth.Use(middlewares.Authenticate)
	rgAuth.POST("/events", PostEvent)
	rgAuth.PUT("/events/:id", UpdateEvent)
	rgAuth.DELETE("/events/:id", DeleteEvent)
	rgAuth.POST("/events/:id/register", RegisterEvent)
	rgAuth.DELETE("/events/:id/cancel", CancelEvent)

	r.POST("/signup", Signup)
	r.POST("/login", Login)

}
