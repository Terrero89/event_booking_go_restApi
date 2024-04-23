package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.Run(":8080") //localhost:8080

}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"events": "Hellow"})
}
