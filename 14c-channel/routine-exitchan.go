package main

import (
	"fmt"
	"time"
)

//goroutine和channel协同工作, 管道是引用类型
//主线程需要等到write和read协程都完成才能退出？读完就写一个标志，主线程读到标志就退出，java的CounterLatch
//如果只写不读，爆满后会包错deadlock, 但是写快读慢爆满，写协程会阻塞不会出错

//var (flag bool = false)

func main() {
	fmt.Println("---------------")

	intChan := make(chan int, 10)  //创建一个管道,如果超过容量，就会阻塞包deadlock,慢慢读也没问题
	exitChan := make(chan bool, 1) //退出管道,而不是标志位,如果要等待n个协程完成呢？

	go writeData(intChan) //启动一个协程
	go readData(intChan, exitChan) //

	for { //无限循环,如果不关闭管道读完还继续读就会抛deadlock,
		//if flag {break} //改为管道通信
		//if v := <- exitChan; v {break}//for range必须关闭管道
		_, ok := <- exitChan //_忽略数据
		if !ok { //读错误，主线程退出
			break
		} //如果读不到东西，读取失败就退出
	}

}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("writeData v=%v\n", i)
		//time.Sleep(time.Millisecond * 10)
	}
	close(intChan) //写完就关闭通道，读依然可以
}

func readData(intChan chan int, exitChan chan bool) {
	//for i := 0; i < 50; i++ {
	//for v := range intChan{ //管道关闭后才能range遍历
	for {
		time.Sleep(time.Millisecond * 100)
		v, ok := <-intChan //关闭不影响读，读完才退出
		if ok {
			//v := <- intChan
			fmt.Printf("readData v=%v\n", v)
		} else {
			break
		}
	}
	//写出一个标志位
	//flag = true //改为管道通信
	exitChan <- true
	time.Sleep(time.Second)
	close(exitChan)//写完关闭通道

}

//go run main.go
//go build -o test.exe
