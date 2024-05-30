package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/internal/models/event"
)

func RegisterEvent(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}

	event, err := event.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't fetch event"})
		return
	}

	if err = event.RegisterAtEvent(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't register at event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "registed successfully at event"})

}

func CancelEvent(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}
	event, err := event.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't fetch event"})
		return
	}

	if err = event.CancelRegisterationEvent(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't cancel event registration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "registration cacelled successfully"})
}
