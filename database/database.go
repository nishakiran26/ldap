package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection = options.Client().ApplyURI("mongodb://localhost:27017")
var client, error = mongo.Connect(context.TODO(), connection)

func CreateConnection() {
	fmt.Println("Create Connection function called")
	if error != nil {
		log.Fatal(error)
	}
	// check the connection
	error = client.Ping(context.TODO(), nil)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connected to MongoDB")
}
func CloseConnection() {
	fmt.Println("Close Connection function called")
	error = client.Disconnect(context.TODO())
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection to MongoDB is closed.")
}
