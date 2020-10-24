package main

import (
	"runtime"
	"sync"
	"fmt"
	//"time"
)

var wg = sync.WaitGroup{} //等待其它协程完工
var counter = 0
var m = sync.RWMutex{} //Mutex互斥锁,RWMutext读写互斥锁

func main() {
	fmt.Printf("Threads:%v\n", runtime.GOMAXPROCS(-1)) //默认最大并行执行的数=cpu
	runtime.GOMAXPROCS(10) //为1时就单cpu，这时就没有parallelism
	//var msg string = "Hello"
	//wg.Add(1) //等待一个
	for i:=0;i<10;i++{
		wg.Add(2) //计数器，类似java的CounterLatch，等待几个,每做完计数减一
		
		m.RLock() //读前先加一个读锁
		go sayHello() //读完在里面释放锁
		m.Lock() //操作前先加一个读写锁
		go increment()//操作完在里面释放锁
	}
	wg.Wait() //类似java的join,等全部routine完成后主线程再继续
	// go func(msg string) { //闭包closure
	// 	fmt.Println(msg)
	// 	wg.Done() //类似java的CounterLatch，做完计数减一
	// }(msg)
	// msg = "bye"
	//time.Sleep(100 * time.Millisecond)
	//msg = "bye"
	
}

func sayHello() {
	//m.RLock() //读锁,错误位置
	fmt.Printf("hello #%v\n", counter)
	m.RUnlock() //释放读锁
	wg.Done()
}

func increment(){
	//m.Lock() //上锁,错误位置
	counter++
	m.Unlock() //释放锁
	wg.Done()
}