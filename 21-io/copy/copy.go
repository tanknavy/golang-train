package main

import (
	"io" //Copy函数调用CopyBuffer
	"fmt"
	"os" //File, OpenFIle
	"bufio" //带缓存的io
)

//io包下的Copy函数,需要提供一个reader和writer
func main() {
	fmt.Println("------------------------------")
	src := "e:/tmp/nasa.jpg"
	dst := "e:/tmp/nasa2.jpg"
	_,err := CopyFile(dst, src)
	if err != nil {
		fmt.Printf("copy file err=%v\n", err)
		return
	}
	fmt.Println("拷贝完成...")
}

func CopyFile(dst string, src string)(written int64, err error){
	srcfile, err := os.Open(src)
	
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer srcfile.Close()

	//通过文件对象，获取到Reader
	reader := bufio.NewReader(srcfile)

	//打开dst文件
	dstfile, err := os.OpenFile(dst, os.O_WRONLY | os.O_CREATE, 0666) //写的方式打开，不存在就创建
		if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer dstfile.Close()
	writer := bufio.NewWriter(dstfile)

	//reader和writer都拿到了，copy
	return io.Copy(writer, reader)
}

//go run main.go
//go build -o test.exe //编译成可执行文件
//go test -v //使用testing框架测试_test.go结尾的文件
