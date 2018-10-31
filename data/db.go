package data

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client
var databaseName = "sgdfgo"

func getClient() *mongo.Client {
	if client == nil {
		var err error
		client, err = mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
		if err != nil {
			panic(err)
		}
	}
	return client

}

func db(name string) *mongo.Collection {
	return getClient().Database(databaseName).Collection(name)
}
