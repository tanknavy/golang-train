package main

import (
	"fmt"
	"strconv"
	"time"
)

//routine是更轻量级的线程，golang单机可以起上个routine,
//进程->线程->协程(轻量级线程,独立的栈空间，共享程序堆空间)

//golang将共享的值通过信道channel传递，多个独立执行的线程从不会主动共享
//不要通过共享内存来通信，而是通过通信来共享内存

func main() {

	fmt.Println("--------1.主线程和协程demo----------------")
	go routineDemo(5) //go开启了一个协程，一般是跑一个函数
	//主线程继续，主线程退出协程马上退出
	for i := 0; i < 10; i++ {
		fmt.Println("main: hello world " + strconv.Itoa(i))
		time.Sleep(1 * time.Second) //每隔一秒
	}

	fmt.Println("--------2.素数,并行协程----------------")
	//sushu(100) //主线程中抛
	//go sushu(200) //协程并行,但主线程如何知道协程啥时候完成
	//time.Sleep(1000)//光等待不靠谱
	//使用channel,写管道exitChan, 或者sync.WaitGroup?
	//详情查看routine + channel


}

//判断素数，计算量非常大，考虑使用go的协程，
func sushu(num int) {
	for i := 2; i <= num; i++ { //从2到num开始
		var flag bool = true
		for j := 2; j < i-1; j++ { //除以 2到n-1之间每个数
			if i%j == 0 {
				flag = false //只要有一个能整除，就不能是素数
				break
			}
		}
		if flag {
			fmt.Println(i)
		}

	}
}

//主线程和协程分别每隔1秒输出"hello"，主线程和协程同时跑，主线程退出协程自动退出
func routineDemo(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("hello world " + strconv.Itoa(i))
		time.Sleep(1 * time.Second) //每隔一秒
	}
}
