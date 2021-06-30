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

func GetSchedulesAndCameras(nowDate string, campusId string) ([]model.SchedulesAndCamers, error) {
	var week int
	if nowDate == ".*" { //提供了http接口,
		today = time.Now().Format("2006-01-02")

	}
	week = int(time.Now().Weekday())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-class-schedules")
	pipeline := []bson.M{ //聚合语句
		{
			"$match": bson.M{
				"campus":           bson.M{"$regex": campusId},
				"from":             bson.M{"$lte": today},
				"to":               bson.M{"$gte": today},
				"schedule.weekday": week,
			},
		},
		{
			"$addFields": bson.M{
				"classroom._id": bson.M{"$toString": "$classroom._id"}, //将_id转换为string
				// "classroom.camerass": bson.M{"_id": "classroom.cameras.gateway", "internetIp": "aaa"}, //将_id转换为string
				// "classroom.camerass": bson.M{"$concatArrays": bson.M["$classroom.camerass""," [gatewayMap]}, //将_id转换为string

				// { homework: { $concatArrays: [ "$homework", [ 7 ] ] } }
				"startTime": "$schedule.startTime",
				"endTime":   "$schedule.endTime",
				"weekday":   "$schedule.weekday",
				"date":      today,
				"week":      week,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "meta-cameras",
				"localField":   "classroom._id",
				"foreignField": "classroom",
				"as":           "classroom.cameras",
			},
		},
	}
	if cursor, err = collection.Aggregate(ctx, pipeline); err != nil {
		log.Fatal(err)
		return nil, err
	}
	//延迟关闭游标
	defer cursor.Close(ctx)

	var results []model.SchedulesAndCamers
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(len(results))
	// for _, result := range results {
	// 	fmt.Println('a', result)
	// }
	return results, err
}
