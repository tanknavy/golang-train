package main

import (
	"fmt"
)

//go run main.go
//go build -o test.exe

//管道的基本操作
//管道(chan关键字)，必须make初始化，有类型，线程安全，FIFO,inteface{}表示全部类型(类似java的Object),
/*
intChan := make(chan int,10)
x,ok := <- intChan 判断读状态
for range没有关闭的管道在读完后会报错deadlock;
只写不读爆满时会报错deadlock,
写快读慢爆满时不会报错,再写协程会阻塞；
channel可以不用那么大
*/

func main() {
	fmt.Println("----")
	//1.定义管道
	var intChan chan int
	intChan = make(chan int, 3) //make初始化，可存放3个int类型的管道

	//2.intChan是什么
	fmt.Printf("intChan的值%v, 本身的地址%p \n", intChan, &intChan) //地址

	//3.向管道写数据
	intChan <- 211 //写数据
	num := 985
	intChan <- num //不能超过chan容量

	//4.看看管道的长度和capacity(容量)，管道固定的
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan)) //%v变量值，%T类型，%p指针

	//5.从读取管道数据
	var num2 int
	num2 = <-intChan //读数据
	<-intChan        //读数据但是扔掉了
	//if num3,ok := <- intChan; !ok {} //一般在循环中这样读取并判断管道是否可读，如果不ok就可以退出读取循环
	fmt.Println(num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//6.interface{}类型的channel
	//var allChan chan interface{}
	allChan := make(chan interface{}, 10)

	allChan <- 985
	allChan <- "tom"
	cat := Cat{"黑猫警长", 4}
	allChan <- cat

	<-allChan//抛出一个不予接受
	<-allChan

	thisCat := <-allChan                                  //接受channel中一个数据
	fmt.Printf("thisCat=%T, thisCat=%v\n", thisCat, thisCat) //运行时可以
	//fmt.Println("thisCat=%v", thisCat.name) //编译错误，管道是interface{}类型，编译器认为它没有任何方法，需要assert
	fmt.Println(thisCat.(Cat).name)//使用类型断言


	//7.管道的关闭和遍历：关闭后不能写但是可以读完，for-range遍历(不close读完会报错)，不能用for i len(intChan),因为len是动态的
	intChan <- 100
	intChan <- 200
	intChan <- 300
	//x,ok := <- intChan //读并判断
	close(intChan) //close是内置函数关闭channal,不能写但是可以读，读完就停止该通道
	for i:= range intChan{ //for range遍历管理时，
		fmt.Println(i)
	}



}

type Cat struct {
	name string
	age  int
}
