package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", GetEvents)
	r.GET("/events/:id", GetEvent)
	r.POST("/events", PostEvent)
	r.PUT("/events/:id", UpdateEvent)
}
