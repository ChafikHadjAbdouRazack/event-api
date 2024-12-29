package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, error := models.GetAllEvents()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	event, error := models.GetEventById(eventId)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {

	var event models.Event
	error := context.ShouldBindJSON(&event)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	event.UserID = context.GetInt64("userId")

	error = event.Save()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	_, error = models.GetEventById(eventId)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	var event models.Event
	error = context.ShouldBindJSON(&event)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	event.ID = eventId

	error = event.Update()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

func deleteEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	event, error := models.GetEventById(eventId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	error = event.Delete()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
