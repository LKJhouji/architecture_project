package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tessa_backend/controller"
	"tessa_backend/zlog"
)

func main() {
	g := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	g.Use(cors.New(corsConfig))
	g.POST("/getPath", controller.GetPath)
	g.POST("/inputPath", controller.InputPath)
	if err := g.Run("0.0.0.0:8000"); err != nil {
		zlog.Fatal(err.Error())
	}
}
