package main

import (
	"fmt"
	_ "gobasic/testInit"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const textPath string = "/Users/zhangjie/Desktop/selfCode/gobasic/files"
const FormatString = "2006年01月02日 15:04:05"

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

func ioutiltest() {
	f, _ := os.Open(textPath + "/b/a.txt")
	b, err := ioutil.ReadAll(f)
	if err == nil {
		fmt.Println(string(b))
		fmt.Println(len(b))
	}
}

func timeProgram() {
	t1 := time.Now()
	fmt.Printf("%T\n", t1)

	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2, _ := time.Parse("2006年01月02日", "1999年10月10日")
	fmt.Println(s2)

	fmt.Println(t1.String())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Minute())
	fmt.Println(t1.YearDay())

	// 时间戳
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	fmt.Println(t4.Unix())
	// 当前时间转时间戳
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	// 加 60s 后
	fmt.Println(t1)
	t5 := t1.Add(time.Duration(time.Second * 60))
	// 求时间间隔
	fmt.Println(t5.Sub(t1))
	t6 := time.Now().Unix()
	rand.Seed(t6)
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)

	t7 := time.Unix(t6, 0)
	fmt.Println(t7.Format("2006年1月2日 15:04:05"))

	time.Sleep(time.Duration(randNum) * time.Second)

}

func fileProgram() {
	fileInfo, err := os.Stat("./files/b/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(fileInfo.Sys())
	// 名称
	fmt.Println(fileInfo.Name())
	// 权限
	fmt.Println(fileInfo.Mode())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 是否为文件夹
	fmt.Println(fileInfo.IsDir())
	// 大小，二进制
	fmt.Println(fileInfo.Size())

	basePath, _ := os.Getwd()
	// 只能创建一层
	_ = os.Mkdir(basePath+"/c", os.ModePerm)
	// 可以创建多层
	_ = os.MkdirAll(basePath+"/a/b/c", os.ModePerm)

	// 绝对路径创建
	file, _ := os.Create(basePath + "/c/a.txt")
	fileInfos, _ := file.Stat()
	fmt.Println(fileInfos.ModTime().Format(FormatString))

	// 相对路径创建
	file2, err := os.Create("./c/b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file2.Name())

	file3, err := os.Open("./c/b/txt")
	fmt.Println(file3)

	file4, err := os.OpenFile("./c/a.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file4)
	file4.Write([]byte("说什么呢"))

	// 只能删除空目录和文件
	err = os.Remove("./c")
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	fileProgram()
}
