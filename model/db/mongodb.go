package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"mageBATestCase/util"
	"time"
)

var client *mongo.Client = nil
var ctx context.Context = nil

func GetConnection() (*mongo.Client, context.Context) {
	if client == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI(util.GetEnvVariable("DATABASE_URL")))
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		return client, ctx
	}

	return client, ctx
}
