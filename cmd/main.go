package main

import (
	"context"
	"fmt"
	"log"
	"myapp/constants"
	"myapp/internal/handlers"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(constants.MongoDBURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(ResponseWriter http.ResponseWriter, HttpRequest *http.Request) {
		ResponseWriter.WriteHeader(http.StatusOK)
		ResponseWriter.Write([]byte("API is running"))
	}).Methods("GET")

	router.HandleFunc("/api/auth", func(ResponseWriter http.ResponseWriter, Request *http.Request) {
		handlers.AuthUserHandler(client, ResponseWriter, Request)
	}).Methods("POST")

	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
