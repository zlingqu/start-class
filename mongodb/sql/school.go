package sql

import (
	// "begin_training/mongodb/init"
	mongoInit "begin_training/mongodb/init"

	// mongoInit "begin_training/mongodb/init"
	"begin_training/mongodb/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSchoolsAndCampuses() ([]model.SchoolsAndCampuses, error) {
	//2.选择数据库和表
	// mmmmClient := m.MongoCli
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// collection := mmmmClient.Database(m.MongoDbName).Collection("meta-campuses")
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-campuses")
	pipeline := []bson.M{ //聚合语句
		{
			"$addFields": bson.M{
				"school": bson.M{"$toObjectId": "$school"},
				// "mySchool": "$school",
				"_id": bson.M{"$toString": "$_id"}, //将_id转换为string
			},
		},
		{
			"$lookup": bson.M{
				"localField":   "school",
				"from":         "meta-schools",
				"foreignField": "_id",
				"as":           "my-school",
			},
		},
		{
			"$unwind": "$my-school",
		},
	}
	if cursor, err = collection.Aggregate(ctx, pipeline); err != nil {
		log.Fatal(err)
		return nil, err
	}
	//延迟关闭游标
	defer cursor.Close(ctx)

	var results []model.SchoolsAndCampuses
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(len(results))
	// for _, result := range results {
	// 	fmt.Println('a', result)
	// }
	return results, err
}

func GetSchoolsInfo() ([]model.Schools, error) {
	//2.选择数据库和表
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(mongoInit.MongoDbName).Collection("meta-schools")
	filter := bson.M{}
	if cursor, err = collection.Find(ctx,
		filter,
		options.Find().SetSkip(0),
		options.Find().SetLimit(200)); err != nil {
		log.Fatal(err)
		return nil, err
	}
	//延迟关闭游标
	defer cursor.Close(ctx)

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.Schools
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	return results, err
}
