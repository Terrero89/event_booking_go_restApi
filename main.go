package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi_go_event_booking/db"

	"restApi_go_event_booking/models"
)

func main() {
	db.InitDB() //initiated DB

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	//server running.
	server.Run(":8080") //localhost:8080'

}
func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events..."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating event"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not create events..."})
	}

	//to indicate we created a new event successfully
	c.JSON(http.StatusCreated, gin.H{"Message": "Event Created", "Event": event})
}
