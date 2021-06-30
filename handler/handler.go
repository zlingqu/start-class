package handler

import (
	"net/http"

	"begin_training/mongodb/sql"
	svc "begin_training/service"

	"github.com/gin-gonic/gin"
)

func GetSchoolInfo(c *gin.Context) {
	data, err := sql.GetSchoolsInfo()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(data),
			"res":   "OK",
			"msg":   "获取成功",
			"data":  data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetSchoolsAndCampuses(c *gin.Context) {
	data, err := sql.GetSchoolsAndCampuses()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(data),
			"res":   "OK",
			"msg":   "获取成功",
			"data":  data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetSchedulesAndCameras(c *gin.Context) {
	data, err := sql.GetSchedulesAndCameras(".*", ".*")

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(data),
			"res":   "OK",
			"msg":   "获取成功",
			"data":  data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetAdjustments(c *gin.Context) {
	data, err := sql.GetAdjustments(".*")

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(data),
			"res":   "OK",
			"msg":   "获取成功",
			"data":  data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetManuleStartClass(c *gin.Context) {
	// data, err := sql.GetClass()
	data, err := svc.GetManuleStartClass()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			// "count": len(data),
			"res":  "OK",
			"msg":  "获取成功",
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetGateway(c *gin.Context) {
	data, err := sql.GetGateway()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			// "count": len(data),
			"res":  "OK",
			"msg":  "获取成功",
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}

func GetTodayClass(c *gin.Context) {
	data, err := sql.GetTodayClass(".*", ".*")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			// "count": len(data),
			"res":  "OK",
			"msg":  "获取成功",
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"count": 0,
		"res":   "error",
		"msg":   "查询数据库失败",
		"data":  err,
	})
}
