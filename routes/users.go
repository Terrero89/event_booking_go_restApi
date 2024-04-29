package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi_go_event_booking/models"
)

func signup(context *gin.Context) {
	var user models.User                 //user unstance
	err := context.ShouldBindJSON(&user) //user being added
	//if there is an error, then show error 400
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	//using new instance to add it to users struct list
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user in Database."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "USER created successfully."})
}
