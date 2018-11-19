package main

import (
	"time"

	"github.com/HAL-Future-Creation-Exhibition/bcp-server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// web
	r.Use(cors)
	r.Static("/web/public", "./docker/web/html")
	r.GET("/web/alive", controller.Web.Alive)
	r.GET("/web/up", controller.Web.Up)
	r.GET("/web/down", controller.Web.Down)
	r.POST("/web/upload", controller.Web.Upload)

	// file
	r.GET("/file/alive", controller.File.Alive)
	r.GET("/file/up", controller.File.Up)
	r.GET("/file/down", controller.File.Down)

	r.GET("/file", controller.File.Ls)
	r.POST("/download/file", controller.File.FileDownload)
	r.POST("/download/dir", controller.File.DirDownload)
	r.POST("/upload/file", controller.File.FileUpload)
	r.POST("/upload/dir", controller.File.DirUpload)
	r.POST("/create/dir", controller.File.CreateDir)
	r.DELETE("/delete/file", controller.File.DeleteFile)
	r.DELETE("/delete/dir", controller.File.DeleteDir)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func cors(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()

}
