package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tessa_backend/constants"
	"tessa_backend/dto/request"
	"tessa_backend/service"
	"tessa_backend/zlog"
)

func GetPath(c *gin.Context) {
	var req request.GetPathRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	res := service.GetRandomPath(req.Index)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func InputPath(c *gin.Context) {
	var req request.InputPathRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	res := service.InputPath(req.Path)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}
