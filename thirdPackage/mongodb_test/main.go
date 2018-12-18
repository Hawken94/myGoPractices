package main

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/davecgh/go-spew/spew"

	"github.com/mongodb/mongo-go-driver/mongo/readpref"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const mongoURL = "mongodb://127.0.0.1:27017"

func main() {
	client, err := mongo.NewClient(mongoURL)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("tmp").Collection("col")
	filter := bson.M{"lovers": 10}
	type result struct {
		Title       string
		Description string
		By          string
		URL         string
		Tags        []int
		Lovers      int
	}
	res := &result{}
	err = collection.FindOne(ctx, filter).Decode(res)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(res)

}
