package repository

import (
	"log"

	"go.mongodb.org/mongo-driver"
)

func NewRepository(url string) Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI(“<<MongoDB Connection URI>>))
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

