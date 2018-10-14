package controller

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"fmt"
)

var Bcp = bcpController{}

type bcpController struct {}

func (b *bcpController) Alive(c *gin.Context) {
	if alive() {
		c.JSON(200, "alive!!")
		return
	}
	c.JSON(200, "no alive...")
}

func alive() bool {
	out, err := exec.Command("sh", "-c", "docker ps | grep bcp-nginx").Output()
	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		return false
	}

	if string(out) != "" {
		return true
	}
	return false
}

func (b *bcpController) Connect(c *gin.Context) {
	if !alive() {
		_, err := exec.Command("make", "run").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}

func (b *bcpController) Disconnect(c *gin.Context) {
	if alive() {
		_, err := exec.Command("make", "stop").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}


