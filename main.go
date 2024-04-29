package main

import (
	"github.com/gin-gonic/gin"
	"restApi_go_event_booking/db"
	"restApi_go_event_booking/routes"
)

func main() {
	db.InitDB() //initiated DB

	server := gin.Default()
	routes.RegisterRoutes(server)

	//server running.
	server.Run(":8080") //localhost:8080'

}
