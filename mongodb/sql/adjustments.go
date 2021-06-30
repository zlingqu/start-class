package sql

import (
	// "begin_training/mongodb/init"
	mongoInit "begin_training/mongodb/init"
	"begin_training/mongodb/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAdjustments(campusId string) ([]model.Adjustments, error) {
	// now := time.Now()
	//2.选择数据库和表
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-adjustments")
	pipeline := []bson.M{ //聚合语句
		{
			"$match": bson.M{
				"campus": bson.M{"$regex": campusId},
				"$or": []bson.M{
					// {"daysOff": "2021-05-21"}, //now.Format("2006-01-02")
					// {"makeUpLessons.at": "2021-05-28"},
					{"daysOff": today},
					{"makeUpLessons.at": today},
				},
			},
		},
	}
	if cursor, err = collection.Aggregate(ctx, pipeline); err != nil {
		log.Fatal(err)
		return nil, err
	}
	//延迟关闭游标
	defer cursor.Close(ctx)

	var results []model.Adjustments
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	log.Println("查到调课记录条数：", len(results))
	for _, result := range results {
		fmt.Println('a', result)
	}
	return results, err
}
