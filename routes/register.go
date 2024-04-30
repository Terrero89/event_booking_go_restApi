package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi_go_event_booking/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered successfully!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	//if err != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
	//	return
	//}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel/unregister event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration cancelled successfully!"})
}
