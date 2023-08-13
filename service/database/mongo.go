package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo(timeout time.Duration) *mongo.Database {
	url := os.Getenv("MOGNO_DB_URI")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return client.Database(os.Getenv("MONGO_DB_NAME"))
}
