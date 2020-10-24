package main

import (
	"fmt"
	"sync"
)

//并发竞争：1)全局锁，2）channel(chan是keyword,通道除了读写数据，还可以作为exitChan标识)
//从共享内存通信，改为信道中分享内存，channel同时只有一个协程能访问
//channel信道, goroutine通信
//channel默认无缓冲，一次仅能处理一个，现在好像可以缓冲了吧，
var wg = sync.WaitGroup{}//等待全部协程完成

func main() {
	//var ch chan int //声明,必须make初始化后才能使用，chan是关键字
	//ch := make(chan int) //信道默认大小1
	ch := make(chan int, 50) //有缓冲信道，如果想放任意类型，声明为interface{}

	//for j := 0; j < 5; j++ {
		wg.Add(2) //类似CountdownLatch, 主线程需要等待两个协程完成才会退出
		//go func() {
		go func(ch <- chan int) { //匿名函数，receive-only,接受一个chan int类型的输入，chan是关键字
			for i := range ch{ //range遍历管道需要管道先close
				//i := <- ch //从channel中接受数据到i
				fmt.Println(i)
			}
			//ch <- 55 //收了再发,错误
			// i = <- ch
			// fmt.Println(i)
			wg.Done() //协程遍历完通道就发送done信号
		}(ch) //匿名函数使用，ch是引用类型，两端操作同一个管道
		
		//go func() {
		go func(ch chan <- int) { //匿名函数，send-only，我发送int类型数据到ch chan通道引用
			i := 29
			ch <- i //向channel中发送数据，默认无缓冲ch,一次一个
			//fmt.Println(<-ch) //收一个并答应
			i = 36
			ch <- i
			close(ch) //如果接收端在循环接受，这里发完了要close(ch)
			wg.Done() //协程遍历完通道就发送done信号
		}(ch)//匿名函数使用，ch是引用类型，两端操作同一个管道
	//}
	wg.Wait() //主线程等待
}
