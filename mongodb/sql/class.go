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

func GetTodayClass(today, campusID string) ([]model.Classes, error) {
	// if today == ".*" { //提供了http接口,
	// 	today = time.Now().Format("2006-01-02")
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-classes")
	filter := []bson.M{
		{
			"$match": bson.M{
				"campus":               bson.M{"$regex": campusID},
				"from":                 bson.M{"$lte": today},
				"to":                   bson.M{"$gte": today},
				"earlyTerminationDate": bson.M{"$exists": true},
			},
		},
	}
	cursor, err = collection.Aggregate(ctx, filter)
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//延迟关闭游标

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.Classes
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	return results, err
}
