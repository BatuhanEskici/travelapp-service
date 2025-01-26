package services

import (
	"context"
	"errors"
	"myapp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthUser(client *mongo.Client, email, password string) (*models.User, error) {
	collection := client.Database("batuhan-apps").Collection("users")

	var user models.User
	err := collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
