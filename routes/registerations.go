package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func registerUser(context *gin.Context) {

	fmt.Println("here")

	userId := context.GetInt64("userId")

	eId := context.Param("id")

	eventId, err := strconv.ParseInt(eId, 10, 64)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "Couldnt Parse Event Id."})
		return
	}

	event, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "Couldnt Fetch Event With That Id."})
		return
	}

	id, err := event.Register(userId)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registeration Id: " + strconv.FormatInt(id, 10)})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eId := context.Param("id")
	id, err := strconv.ParseInt(eId, 10, 64)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "Couldnt Parse Event Id."})
		return
	}

	event, err := models.GetSingleEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Couldnt Find Event With That Id."})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration Cancelled Successfully."})

}
