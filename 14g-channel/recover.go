package main

import (
	"time"
	"fmt"
)

//使用recover阻止其它协程panic而导致整个程序崩溃
func main() {
	fmt.Println("---------------")
	go sayHi() //ok
	go test() //error,不想因为这个协程失败而导致大家都玩完，函数里面defer+recover捕获异常
	//time.Sleep(time.Second) 
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("main() ok=%d\n",i)
	}
}

func sayHi(){
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("routine() hello,world")
	}
}

func test(){//defer + recover
	defer func(){ //匿名函数
		//捕获抛出的panic
		if err := recover(); err != nil {//如果发现错误
			fmt.Println("test发生错误",err) //处理错误
		}
	}() //匿名函数调用一下
	var myMap map[int]string
	myMap[0] ="golang" //error,没有初始化就赋值

}
//go run main.go
//go build -o test.exe
