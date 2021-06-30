package sql

import (
	mongoInit "begin_training/mongodb/init"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	err        error
	cursor     *mongo.Cursor
)

var today = time.Now().Format("2006-01-02")

var mongoClient *mongo.Client = mongoInit.MongoCli
