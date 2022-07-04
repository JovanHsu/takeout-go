package file

import (
	"io/ioutil"
	"log"
)

func FindResource(filename string) bool {
	return listFiles(filename, ".", 0)
}

func listFiles(searchFile string, dirname string, level int) bool {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		if fi.Name() == searchFile {
			return true
		}
		filename := dirname + "/" + fi.Name()
		if fi.IsDir() {
			//继续遍历fi这个目录
			return listFiles(searchFile, filename, level+1)
		}
	}
	return false
}
