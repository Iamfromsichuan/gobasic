package main

import (
	"fmt"
	"os"
	"strings"
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
	r := strings.NewReader("hello 世界")
	b := make([]byte, r.Size())
	n, err := r.Read(b)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("读取的长度为", n)
	fmt.Println("数据内容为", string(b))
}

func ReaderFile() {
	f, _ := os.Open(textPath + "/b/a.txt")
	fileinfo, _ := f.Stat()
	b := make([]byte, fileinfo.Size())
	n, _ := f.Read(b)
	fmt.Println("读取的长度为", n)
	fmt.Println("数据内容为", string(b))
}

func writeFile() {
	filepath := textPath + "/b/b.txt"
	f, err := os.OpenFile(filepath, os.O_APPEND, 0666)
	if err != nil {
		f, err = os.Create(filepath)
	}
	f.Write([]byte("展望"))
	defer func() {
		f.Close()
	}()
}

func main() {
	ReaderFile()
}
