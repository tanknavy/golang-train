package main

import (
	"fmt"
)

//routine是更轻量级的线程，golang单机可以起上个routine,
//进程->线程->协程(轻量级线程,独立的栈空间，共享程序堆空间)

//golang将共享的值通过信道channel传递，多个独立执行的线程从不会主动共享
//不要通过共享内存来通信，而是通过通信来共享内存

func main() {

	fmt.Println("--------素数,并行协程----------------")
	//sushu(100) //主线程中抛
	//go sushu(200) //协程并行,但主线程如何知道协程啥时候完成
	//time.Sleep(1000)//光等待不靠谱
	//使用channel,写管道exitChan, 或者sync.WaitGroup?
	//详情查看routine + channel
	num := 1000 //最大计算到几个素数
	intChan := make(chan int, num)   //1-n的数字，1个routine
	primeChan := make(chan int, num) //素数管道, 2个routine
	exitChan := make(chan bool, 4)   //全部协程完成标志管道
	
	go pushData(num, intChan, exitChan) //一个协程向intChan放入数字
	
	//起多少个协程
	totalRoutines := 4 //要起/等几个协程完工？和cpu一样多

	for i := 0; i < totalRoutines; i++ {
		go primeData(intChan, primeChan, exitChan) //从intChan取出数据并判断是否为素数，如果是放入primeChan
	}

	//主线程等待全部协程完工
	//方法一
	// count := 0 //当前几个协程完工，还可以使用sync.WaitGroup
	// for count < totalRoutines { //主线程等待协程完工
	// 	_, ok := <-exitChan
	// 	if ok {
	// 		count ++ 
	// 	} else {
	// 		break
	// 	}
	// }
	
	//方法二：从exitChan能读到几次数据就可以退出,推荐使用sync.WaitGroup
	go func(){ //匿名函数，起一个协程看看其它协程完工没有
		for i := 0; i < totalRoutines; i++ { //只能读几次，多了会有报deadlock
			<- exitChan //能得到几次完工标志就表示ok
		}
		//记得primeChan和exitChan还没关闭
		close(primeChan) //全部写primeChan完工就关闭
		close(exitChan)
	}() //匿名函数启动
	
	//遍历取出数据
	fmt.Println("开始读取primeChan全部素数")
	for {
		v, ok := <- primeChan //一直取，知道全部读完
		if !ok {
			break
		}else{
			fmt.Println(v)
		}
	}
	fmt.Println("主线程退出")

}

//给定一个数字n，输出1~n之间全部素数，
func sushu(num int) {
	for i := 2; i <= num; i++ { //从2到num开始
		var flag bool = true
		for j := 2; j < i; j++ { //除以 2到n-1之间每个数
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

//给定一个数字n, 判断它是否是素数
func isPrime(i int) bool {
	if(i==1) {return true}
	if(i==2) {return false}
	var flag bool = true
	for j := 2; j < i; j++ { //除以 2到n-1之间每个数
		if i%j == 0 {
			flag = false //只要有一个能整除，就不能是素数
			break //退出forx循环
		}
	}
	//if flag {fmt.Println(i)}
	return flag
}

//写数字到管道，等待读取判读是否素数
func pushData(num int, intChan chan int, exitChan chan bool) {
	for i := 1; i <= num; i++ {
		intChan <- i
	}
	close(intChan) //写数字通道完关闭
	//exitChan <- true //做完标志位
}

//从数字管道中拿到数字，判断是素数就写入素数管道
func primeData(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok { //读管道异常，没有数据
			break //不再读取管道
		} else {
			if isPrime(v) {
				primeChan <- v
			}
		}
	}

	fmt.Println("一个协程因为取不到数据而退出")
	exitChan <- true //本协程完工标志
	//close(primeChan) //多个协程情况下谁来关闭素数管道？
	//close(exitChan)  //多个协程情况下谁来关闭退出管道？
}
