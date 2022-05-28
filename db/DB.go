package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"piterdev.com/app/models"
)

var ctx = context.TODO()
var client = ConnectDB()
var userColl = client.Database("main").Collection("users")
var calendarColl = client.Database("main").Collection("calendar")
var eventColl = client.Database("main").Collection("events")

type User *models.User

func ConnectDB() *mongo.Client {

	godotenv.Load()

	var dbUser string = os.Getenv("MONGO_USER")
	var dbPass string = os.Getenv("MONGO_PSW")
	var dbDm string = os.Getenv("MONGO_DM")

	var uri string = "mongodb+srv://" + dbUser + ":" + dbPass + "@" + dbDm + "/?retryWrites=true&w=majority"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
