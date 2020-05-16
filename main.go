package main

import (
	"fmt"
	_ "gobasic/testInit"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const textPath string = "/Users/zhangjie/Desktop/selfCode/gobasic/files"
const FormatString = "2006年01月02日 15:04:05"

var RootPath, _ = os.Getwd()

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

	// 只能删除空目录和文件 rm
	err = os.Remove("./c")
	if err != nil {
		fmt.Println(err)
	}
	//  删除全部 rm -rf
	// os.RemoveAll("")

}

func ioProgram() {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile(basePath+"/files/b/a.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	bs := make([]byte, 32*1024)
	var c []byte
	var total = 0
	for {
		n, err := file.Read(bs)
		total = total + n
		if err != nil && err == io.EOF {
			fmt.Println(n)
			fmt.Println(err)
			break
		}
		if n == 0 {
			break
		}
		c = append(c, bs[:n]...)
		fmt.Println(string(c))
	}
	fmt.Println(total)
}

func ioWrite() {
	basePath, _ := os.Getwd()
	file, err := os.OpenFile(basePath+"/c/c.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte("motorola"))
}

func seekTest() {
	fileName := RootPath + "/c/a.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs := []byte{0}
	file.Read(bs)

	// 从开始位置向后偏移4个字节
	file.Seek(4, io.SeekStart)
	file.Read(bs)

	// 从当前光标向后偏移两个位置
	file.Seek(2, io.SeekCurrent)
	file.Read(bs)

	// 将光标移动到末尾
	file.Seek(0, io.SeekEnd)
	file.Write([]byte("就是最后咯"))
	fmt.Println(string(bs))
}

//  c -- 短点续传
func breakPointerContinue() {
	fileName := RootPath + "/files/b/a.txt"
	name := fileName[strings.LastIndex(fileName, "/")+1:]
	temp := name + "temp.txt"
	fmt.Println(name, temp)

	file1, err := os.Open(fileName)
	handleError(err)
	file2, err := os.OpenFile(RootPath+"/axc.txt", os.O_CREATE|os.O_RDWR, 0777)
	handleError(err)
	file3, err := os.OpenFile(RootPath+"/"+temp, os.O_CREATE|os.O_RDWR, 0777)
	handleError(err)

	defer file2.Close()
	defer file1.Close()

	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	// handleError(err)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	// handleError(err)
	fmt.Println(count)

	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)

	data := make([]byte, 1024, 1024)
	n2 := -1
	n3 := -1
	total := int(count)

	for {
		n2, err = file1.Read(data)
		if err != nil && err == io.EOF {
			fmt.Println("读取完毕", total)
			file3.Close()
			os.Remove(temp)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Println(total)
		//if total > 8000 {
		//	panic("假装错误")
		//}
	}
}

func chanel() {
	ch1 := make(chan<- int)
	//ch2 := make(<- chan int)
	// ch1 <- 112
	data := <-ch1
	//fmt.Println(data)
	//ch2 <- 22
}

func senData(chi chan string, done chan bool) {
	chi <- "四川话"
	done <- true
}

func main() {
	chanel()
}

func handleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
