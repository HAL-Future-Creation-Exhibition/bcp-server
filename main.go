package main

import (
	"github.com/gin-gonic/gin"
	"github.com/HAL-Future-Creation-Exhibition/bcp-server/controller"
)

func main() {
	r := gin.Default()
	r.GET("/alive", controller.Bcp.Alive)
	r.GET("/connect", controller.Bcp.Connect)
	r.GET("/disconnect", controller.Bcp.Disconnect)
	r.Run() // listen and serve on 0.0.0.0:8080
}
