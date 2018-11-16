package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	"github.com/HAL-Future-Creation-Exhibition/bcp-server/util"
	"github.com/gin-gonic/gin"
)

var File = fileController{"./docker/file/html/tmp/"}

type fileController struct {
	Base string
}

func (f *fileController) Ls(c *gin.Context) {
	if !f.alive() {
		c.JSON(404, "ストレージが有効化されてません。")
		return
	}
	path := c.DefaultQuery("path", "")

	type Raw struct {
		CurrentPath string `json:"current_path"`
		Name        string `json:"name"`
		IsDir       bool   `json:"isDir"`
	}

	type Res struct {
		Raws []Raw `json:"raws"`
	}

	fis, err := ioutil.ReadDir(f.Base + path)

	if err != nil {
		c.JSON(404, gin.H{"message": "ディレクトリが存在しません。"})
		return
	}

	var res Res
	for _, info := range fis {
		fmt.Println(info)
		res.Raws = append(res.Raws, Raw{
			CurrentPath: path + "/",
			Name:        info.Name(),
			IsDir:       info.IsDir(),
		})
	}

	c.JSON(200, res)
}

func (f *fileController) FileUpload(c *gin.Context) {

}

func (f *fileController) DirUpload(c *gin.Context) {

}

func (f *fileController) Download(c *gin.Context) {
	if !f.alive() {
		c.JSON(404, "ストレージが有効化されてません。")
		return
	}
	req := struct {
		Paths []string
	}{}
	c.BindJSON(&req)
	fmt.Println(req.Paths)
	if len(req.Paths) == 0 {
		c.JSON(http.StatusBadRequest, "ファイル、ディレクトリが指定されていません。")
		return
	}

	header := c.Writer.Header()
	header["Content-Type"] = []string{"application/octet-stream"}
	if len(req.Paths) == 1 {
		header["Content-Disposition"] = []string{"attachment; filename=" + req.Paths[0]}
		c.File(f.Base + req.Paths[0])
		return
	}

	fileName := "bcp-download-" + time.Now().Format("2006-01-02") + ".zip"
	output := "zip/" + fileName

	util.File.TransZip(f.Base, output, req.Paths)

	header["Content-Disposition"] = []string{"attachment; filename=" + fileName}
	c.File(output)
}

func (f *fileController) Alive(c *gin.Context) {
	if f.alive() {
		c.JSON(200, "alive!!")
		return
	}
	c.JSON(200, "no alive...")
}

func (f *fileController) alive() bool {
	out, err := exec.Command("sh", "-c", "docker ps | grep bcp-file").Output()

	if err != nil {
		fmt.Println(err)
		return false
	}

	if string(out) != "" {
		return true
	}
	return false
}

func (f *fileController) Up(c *gin.Context) {
	if !f.alive() {
		_, err := exec.Command("make", "file/up").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}

func (f *fileController) Down(c *gin.Context) {
	if f.alive() {
		_, err := exec.Command("make", "file/down").Output()

		if err != nil {
			c.JSON(500, "重大エラー")
			return
		}
	}
	c.JSON(200, "ok!")
}