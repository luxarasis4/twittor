package database

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username = "lasis"
var password = "La1124044454#"
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + url.QueryEscape(username) + ":" + url.QueryEscape(password) + "@cluster0.itc6b9e.mongodb.net/?retryWrites=true&w=majority")
var MongoCN = connection()

func connection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection to database was successfully")

	return client
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
