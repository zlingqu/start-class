package sql

import (
	// "begin_training/mongodb/init"
	mongoInit "begin_training/mongodb/init"
	"begin_training/mongodb/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGateway() ([]model.Gateway, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-gateways")
	filter := bson.M{}
	if cursor, err = collection.Find(ctx, filter); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []model.Gateway
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}
