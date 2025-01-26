package handlers

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/services"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func AuthUserHandler(client *mongo.Client, ResponseWriter http.ResponseWriter, Request *http.Request) {
	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(Request.Body).Decode(&requestData)

	if err != nil {
		http.Error(ResponseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.AuthUser(client, requestData.Email, requestData.Password)

	if err != nil {
		http.Error(ResponseWriter, err.Error(), http.StatusUnauthorized)
		return
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ResponseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ResponseWriter).Encode(userResponse)
}
