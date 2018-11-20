package util

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
)

var File fileUtil

type fileUtil struct{}

func (f *fileUtil) zipAddFile(w *zip.Writer, workDir, name string) {
	dat, err := ioutil.ReadFile(workDir + "/" + name)
	if err != nil {
		panic(err)
	}

	fi, err := w.Create(name)
	if err != nil {
		panic(err)
	}
	_, err = fi.Write(dat)
	if err != nil {
		panic(err)
	}
}

func (f *fileUtil) ZipAddDir(w *zip.Writer, workDir, path string) {
	fis, err := ioutil.ReadDir(workDir + path)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		if fi.IsDir() {
			//new := jun + "/" +  hoge
			newWorkDir := path + "/" + fi.Name()
			fmt.Println("Recursing and Adding SubDir: " + fi.Name())
			fmt.Println("Recursing and Adding SubDir: " + newWorkDir)
			f.ZipAddDir(w, workDir, newWorkDir)
		} else {
			f.zipAddFile(w, workDir, path+"/"+fi.Name())
		}
	}
}
