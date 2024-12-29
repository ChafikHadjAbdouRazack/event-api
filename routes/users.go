package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
     var user models.User
	 error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	error = user.ValidateCredentials()

	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": error.Error()})
		return
	}
	jwtToken, error := utils.GenerateToken(user.Email,user.ID)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})
}
func signUp(context *gin.Context) {
	var user models.User
	error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	error = user.Save()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
