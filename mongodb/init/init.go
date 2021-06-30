package mongoInit

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCli    *mongo.Client
	MongoUri    string
	MongoDbName string
)

func init() {
	// MongoUri = "mongodb://admin:admin@10.12.19.5:31001/admin?authSource=admin"
	// MongoDbName = "lexue-dev"
	// var err error
	MongoUri = "mongodb://admin:admin@192.168.3.141:31001/admin?authSource=admin"
	MongoDbName = "lexue-test"

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions := options.
		Client().
		SetConnectTimeout(time.Second * 3).
		ApplyURI(MongoUri)

	var err error
	MongoCli, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("无法连接mongodb,启动失败: ", err)
		panic(err)
	}
	// 检查连接
	err = MongoCli.Ping(ctx, nil)
	if err != nil {
		log.Fatal("无法连接mongodb,启动失败: ", err)
		panic(err)
	}
}
