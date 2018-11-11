package util

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

var File fileUtil

type fileUtil struct{}

func (f *fileUtil) TransZip(baseFolder, outFilePath string, filePaths []string) {
	outFile, err := os.Create(outFilePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	f.zipAdd(w, baseFolder, "", filePaths)

	if err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}
}

func (f *fileUtil) zipAdd(w *zip.Writer, baseFolder, baseInZip string, filePaths []string) {
	for _, path := range filePaths {
		file, err := os.Open(baseFolder + path)

		if err != nil {
			panic(err)
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if fi.IsDir() {
			newBase := baseFolder + fi.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + fi.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)
			f.zipAddDir(w, newBase, fi.Name(), path)
		} else {
			f.zipAddFile(w, baseFolder, baseInZip, fi.Name())
		}
	}
}

func (f *fileUtil) zipAddFile(w *zip.Writer, basePath, baseInZip, name string) {
	dat, err := ioutil.ReadFile(basePath + name)
	if err != nil {
		panic(err)
	}

	fmt.Println(baseInZip)
	fi, err := w.Create(baseInZip + "/" + name)
	if err != nil {
		panic(err)
	}
	_, err = fi.Write(dat)
	if err != nil {
		panic(err)
	}
}

func (f *fileUtil) zipAddDir(w *zip.Writer, basePath, baseInZip, path string) {
	fis, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	var filePaths []string
	for _, info := range fis {
		filePaths = append(filePaths, info.Name())
	}
	fmt.Println("files")
	fmt.Println(filePaths)
	f.zipAdd(w, basePath, baseInZip, filePaths)
}
