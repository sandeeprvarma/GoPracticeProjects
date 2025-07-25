package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"message": "Hello Sandeep!"})
	context.JSON(http.StatusOK, event)
}

func getAllEvents(context *gin.Context) {
	// fmt.Println("Events")
	// var newEvents = models.Events{
	// 	ID:         1,
	// 	Title:      "ABCD",
	// 	UserId:     123,
	// 	Created_At: time.Now(),
	// }
	// newEvents.Save()
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
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
	events.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully!"})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	_, err = models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not a valid Id"})
		return
	}

	var updatedEvents models.Events
	err = context.ShouldBindJSON(&updatedEvents)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	updatedEvents.UpdateEvent(id)
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully!"})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	_, err = models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not a valid Id"})
		return
	}

	var updatedEvents models.Events
	updatedEvents.Delete(id)
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully!"})

}
