package router

import (
	"net/http"

	"begin_training/handler"

	"github.com/gin-gonic/gin"
)

// type api string

const (
	ApiV1 = "/api/v1"
	ApiV2 = "/api/v2"
)

var (
	isHealthCHeck = true
	routerGroups  map[string]*gin.RouterGroup
)

func Init(engine *gin.Engine) *gin.Engine {

	routerGroups = map[string]*gin.RouterGroup{ //map实例化
		ApiV1: engine.Group(ApiV1), //定义多个routeGroup,group包含了基础路径(basePath)
		ApiV2: engine.Group(ApiV2),
	}
	httpDefaultConfig(engine) //http默认的路由等配置
	// engine.Use(gin.Logger())
	v1 := routerGroups[ApiV1]
	v2 := routerGroups[ApiV2]
	deployRoute(v1, v2) //添加路由

	return engine

}

func healCheck(c *gin.Context) {
	c.String(200, "# TYPE health_info gauge\n"+"health_info{status=\"ok\", name=\"start-class\", namespace=\"xmc2-lexue\"} 0\n")
}

func httpDefaultConfig(engine *gin.Engine) {

	if isHealthCHeck {
		engine.GET("/healthCheck", healCheck)
		isHealthCHeck = true
	}

	engine.NoRoute(func(c *gin.Context) { //自定义404页面
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "StatusNotFound，找不到该路由",
		})
		return
	})

	engine.NoMethod(func(c *gin.Context) { //自定义no method
		c.JSON(http.StatusOK, gin.H{
			"code": 405,
			"msg":  "StatusMethodNotAllowed，找不到该方法",
		})
		return
	})

}

func deployRoute(v1 *gin.RouterGroup, v2 *gin.RouterGroup) {
	v1.GET("/healthCheck", healCheck)
	v1.GET("/schools", handler.GetSchoolInfo)
	v1.GET("/schoolsAndCampuses", handler.GetSchoolsAndCampuses)
	v1.GET("/schedulesAndCameras", handler.GetSchedulesAndCameras)
	v1.GET("/getAdjustments", handler.GetAdjustments)
	v1.GET("/getGateways", handler.GetGateway)
	v1.GET("/getTodayClass", handler.GetTodayClass)
	v1.GET("/startClass", handler.GetManuleStartClass)
	v2.GET("/healthCheck", healCheck)
}
