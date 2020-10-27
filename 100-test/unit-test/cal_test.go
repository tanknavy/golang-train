//package main //测试中main包也可以
package myTest

import (
	"fmt" //前面加_暂时注销
	//go的轻量级测试框架testing，通过go test命令，自动执行func TestXxx(*testing.T)，比如TestAdd(),T是结构体
	"testing" 
)

// 命名规范：文件名test.go结尾，这里函数TestXxx
//编写一个测试用例，测试cal.go文件中的函数
func TestAddUpper(t *testing.T){ //命名必须是Test+被测试函数名
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误，期望=%d,实际=%d\n", 55, res) //传统方法
		t.Fatalf("AddUpper(10)执行错误，期望=%d,实际=%d\n", 55, res) //相当于调用Logf之后调用FailNow， 就是输出log并停止程序
	}

	//如果正确
	t.Logf("AddUpper(10)执行ok")

	res2 := sub(5,3)
	if res2 != 2 {
		//fmt.Printf("AddUpper(10)执行错误，期望=%d,实际=%d\n", 55, res) //传统方法
		t.Fatalf("sub(5,3)执行错误，期望=%d,实际=%d\n", 2, res) //相当于调用Logf之后调用FailNow， 就是输出log并停止程序
	}

	//如果正确，输出日志
	t.Logf("sub(5,3)执行ok")
}

func TestHello(t *testing.T){
	fmt.Println("TestHello函数被调用")
}

//go run main.go
//go build -o test.exe
//go test 测试_test.go结尾的文件