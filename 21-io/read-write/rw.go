package main

import (
	"fmt"
	"io/ioutil" //一次操作全部
)

func main() {
	fmt.Println("------------------------------")
	//需求:将文件1的内容复制到文件2
	//1.将file1内容读到内存
	f1 := "e:/tmp/t01.txt"
	f2 := "e:/tmp/t03.txt"

	data, err := ioutil.ReadFile(f1)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(f2, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v", err)
		return
	}

}

//go run main.go
//go build -o test.exe //编译成可执行文件
//go test -v //使用testing框架测试_test.go结尾的文件
