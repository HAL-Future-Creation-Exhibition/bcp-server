package main

import (
	"github.com/HAL-Future-Creation-Exhibition/bcp-server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// web
	r.GET("/web/alive", controller.Web.Alive)
	r.GET("/web/up", controller.Web.Up)
	r.GET("/web/down", controller.Web.Down)

	// file
	r.GET("/file/alive", controller.File.Alive)
	r.GET("/file/up", controller.File.Up)
	r.GET("/file/down", controller.File.Down)

	r.GET("/file", controller.File.Ls)
	r.POST("/download", controller.File.Download)
	r.Run() // listen and serve on 0.0.0.0:8080
}
