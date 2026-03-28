package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})

}

func createEvent(context *gin.Context) {

	userId := context.GetInt64("userId")

	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}

	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message2": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created Event.", "event": event})
}

func getEventById(context *gin.Context) {

	fmt.Println("GR1")
	eventId := context.Param("id")
	fmt.Println("Id", eventId)

	id, err := strconv.ParseInt(eventId, 10, 64)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Couldnt Parse The Event Integer Id."})
		return
	}

	event, err := models.GetSingleEvent(id)
	fmt.Println(event)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": event})

}

func updateEvent(context *gin.Context) {
	fmt.Println("u1")

	eventId := context.Param("id")

	id, err := strconv.ParseInt(eventId, 10, 64)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Couldnt Parse Id."})
		return
	}

	_, err = models.GetSingleEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.Id = id

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Update Successfully"})

}

func deleteEvent(context *gin.Context) {

	fmt.Println("Delete")
	eventId := context.Param("id")

	id, err := strconv.ParseInt(eventId, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt Parse Integer Id."})
		return
	}

	event, err := models.GetSingleEvent(id)

	// fmt.Println(event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt Fetch Event."})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt Fetch Event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully"})

}
