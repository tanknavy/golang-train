package main

import (
	"time"
	"fmt"
	"sync"
)

//go run main.go
//go build -o test.exe

//sync包下面的互斥锁
//计算1-200各个数的阶乘, 并发/并行安全问题，判断是否存在资源竞争，编译程序时增加一个参数-race
//并发竞争: 共享的值通过信道传递，一个时间点只有一个go协程能访问，提倡不通过共享内存来通信，而通过通信(信道)来共享内存，
//res := make([]int,20) //类型推导的变量不能放这里，改为var
//var res []int = make([]int,20)
var (
	myMap = make(map[int]uint64, 10) //并发写安全问题go build -race main.go
	//golang中锁解决方案：全局锁(互斥锁，读写锁)，channel
	lock sync.Mutex //互斥锁，包含Lock和Unlock
)

func main() {
	fmt.Println("----")
	//go build -race main.go//测试竞争问题
	//go jiecheng(10) //计算n阶乘
	for i := 0; i < 100; i++ {
		//go build -race main.go然后跑main.exe文件Found 3 data race(s)
		go jiecheng(i+1) //开启多个协程后,fatal error: concurrent map writes,并发写问题
	}
	
	time.Sleep(2 * time.Second)//主线程退出，协程马上退出,等待多久合适，使用channel吧
	lock.Lock() //函数
	fmt.Println(myMap) //可能这里需要锁
	lock.Unlock()
}

func jiecheng(num int){//阶乘计算，结果放入map
	//res := [10]int{} //数组中个数不能是变量
	//list := make([]int,num) //slice
	//myMap := make(map[int]int, num)
	var tmp uint64 = 1
	for i:=1;i<=num;i++{
		//tmp *= i
		tmp += uint64(i) //阶乘太大，改为求累加
		//list[i-1] = tmp
		//myMap[i] = tmp
	}
	//加锁
	lock.Lock() //互斥锁，其它协程排队
	myMap[num] = tmp
	lock.Unlock() //解锁
	//fmt.Println(list)
	//fmt.Println(myMap)
}