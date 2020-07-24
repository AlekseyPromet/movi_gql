package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver@v1.3.5"
)

//NewRepository create connection
func NewRepository(url string) Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return client
}
