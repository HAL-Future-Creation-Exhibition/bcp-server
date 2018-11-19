package controller

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var Web = webController{"./docker/web/html/index.html", "./docker/web/html/index.css", "./docker/web/html/index.js"}

type webController struct {
	HtmlPath string
	CssPath  string
	JsPath   string
}

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

func (w *webController) Upload(c *gin.Context) {
	var req struct {
		Html string
		Css  string
		Js   string
	}

	c.BindJSON(&req)
	if err := os.Remove(w.HtmlPath); err != nil {
		panic(err)
	}
	if err := os.Remove(w.CssPath); err != nil {
		panic(err)
	}
	if err := os.Remove(w.JsPath); err != nil {
		panic(err)
	}
	w.writeFile(w.HtmlPath, req.Html)
	w.writeFile(w.CssPath, req.Css)
	w.writeFile(w.JsPath, req.Js)

}

func (w *webController) writeFile(name, str string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte(str)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
