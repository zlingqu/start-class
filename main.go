package main

import (
	"begin_training/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// c := mongoInit.GetMgoCli()
	// mongoInit.GetMgoCli()

	// k := kafkaInit.InitKafka()
	// p.ProductMessToTopicQuzl(k)
	// sql.GetSchoolsInfo(mongoInit.MongoCli)

	r := gin.Default()
	router.Init(r)
	gin.SetMode(gin.ReleaseMode)
	r.Run(":8080")

}
