package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events/db"
	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})

}

func createEvent(context *gin.Context) {
	// fmt.Println("1")
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}
	// event.Id = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message2": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created Event.", "event": event})
}

func getEventById(context *gin.Context) {
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
