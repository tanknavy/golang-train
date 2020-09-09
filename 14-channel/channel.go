package main

import (
	"fmt"
	"sync"
)

//channel信道, goroutine通信
//channel默认无缓冲，一次仅能处理一个
var wg = sync.WaitGroup{}

func main() {
	//ch := make(chan int) //信道默认大小1
	ch := make(chan int, 50) //有缓冲信道

	//for j := 0; j < 5; j++ {
		wg.Add(2)
		//go func() {
		go func(ch <- chan int) { //receive-only
			for i := range ch{
				//i := <- ch //从channel中接受数据到i
				fmt.Println(i)
			}
			//ch <- 55 //收了再发,错误
			// i = <- ch
			// fmt.Println(i)
			wg.Done()
		}(ch)
		
		//go func() {
		go func(ch chan <- int) { //send-only
			i := 29
			ch <- i //向channel中发送数据，默认无缓冲ch,一次一个
			//fmt.Println(<-ch) //收一个并答应
			i = 36
			ch <- i
			close(ch) //如果接收端在循环接受，这里发完了要close(ch)
			wg.Done()
		}(ch)
	//}
	wg.Wait() //主线程等待
}
