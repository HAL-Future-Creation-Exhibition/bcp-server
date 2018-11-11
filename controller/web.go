package controller

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var Web = webController{}

type webController struct{}

func (w *webController) Alive(c *gin.Context) {
	if w.alive() {
		c.JSON(200, "alive!!")
		return
	}
	c.JSON(200, "no alive...")
}

func (w *webController) alive() bool {
	out, err := exec.Command("sh", "-c", "docker ps | grep bcp-web").Output()
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

func (w *webController) Up(c *gin.Context) {
	if !w.alive() {
		_, err := exec.Command("make", "web/up").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}

func (w *webController) Down(c *gin.Context) {
	if w.alive() {
		_, err := exec.Command("make", "web/down").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}
