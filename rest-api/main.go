package main

// https://pkg.go.dev/

import (
	"net/http"
	"rest-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getAllEvents(context *gin.Context) {
	// fmt.Println("Events")
	var newEvents = models.Events{
		ID:         1,
		Title:      "ABCD",
		UserId:     123,
		Created_At: time.Now(),
	}
	newEvents.Save()
	events := models.GetAllEvents()
	// context.JSON(http.StatusOK, gin.H{"message": "Hello Sandeep!"})
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var events models.Events
	err := context.ShouldBindJSON(&events)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully!"})
}
