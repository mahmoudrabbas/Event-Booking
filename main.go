package main

import (
	"fmt"
	"net/http"

	"example.com/events/db"
	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()

	context.JSON(http.StatusOK, gin.H{"events": events})

}

func createEvent(context *gin.Context) {
	fmt.Println("D")
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	event.Id = 1
	event.UserId = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Created Event.", "event": event})
}
