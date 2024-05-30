package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/internal/models/event"
)

func GetEvents(ctx *gin.Context) {
	events, err := event.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't fetch events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, event)
}

func PostEvent(ctx *gin.Context) {

	var event event.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "unsupported format"})
		return
	}
	event.UserID = ctx.GetInt64("uid")
	if err := event.Save(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid json format"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "event created sucessfully"})
}

func UpdateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}

	existingEvent, err := event.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't fetch event"})
		return
	}

	if existingEvent.UserID != ctx.GetInt64("uid") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthirized operation"})
		return
	}

	var updatedEvent event.Event
	if err = ctx.BindJSON(&updatedEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid json format"})
		return
	}

	updatedEvent.ID = id
	if err = updatedEvent.UpdatEvent(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "event updated successfully"})
}

func DeleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}

	existingEvent, err := event.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't fetch event"})
		return
	}

	if existingEvent.UserID != ctx.GetInt64("uid") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthirized operation"})
		return
	}

	if err = existingEvent.DeleteEvent(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't delete event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "event deleted successfully"})
}
