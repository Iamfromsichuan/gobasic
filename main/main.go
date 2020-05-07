package main

import (
	"fmt"
	"os"
)

const textPath string = "/Users/zhangjie/Desktop/selfCode/gobasic/files"

func createDir() {
	err := os.Mkdir(textPath+"/test", os.ModeDir)
	if err != nil {
		fmt.Println("文件创建失败,失败原因: ", err)
		return
	}
	fmt.Println("文件创建成功")
}

func createDir2() {
	err := os.MkdirAll(textPath+"/b/c", os.ModeDir)
	if err != nil {
		fmt.Println("文件创建失败,失败原因: ", err)
		return
	}
	fmt.Println("文件创建成功")
}

func opentest() {
	f, err := os.Open(textPath + "/b/a.txt")
	if err != nil {
		fmt.Println("文件打开失败,失败原因: ", err)
		return
	}
	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println("文件读取失败,失败原因: ", err)
		return
	}
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.Size())
}

func osRemove() {
	err := os.Remove(textPath + "/b/c")
	if err != nil {
		fmt.Println("文件删除失败,失败原因: ", err)
		return
	}
	fmt.Println("删除成功")
}

func ReaderTest() {

}

func main() {
	ReaderTest()
}
