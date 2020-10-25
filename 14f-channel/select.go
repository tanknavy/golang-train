package main

import (
	"fmt"
	"strconv"
	"time"
)

//管道不close读取时会阻塞，读完后还会报deadlock
//使用select可以解决从管道取数据的阻塞问题
func main() {
	fmt.Println("---------------")

	//1.定义一个int管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个string管道
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}

	//3.传统的方法在遍历管道时，如果不关闭会阻塞导致死锁deadlock,但什么时候关闭很困难
	//在实际开发中，不好确定什么时候付关闭该管道，使用select
	lable:
	for { //无限循环
		select { //select语法
		case v := <-intChan:
			//注意:如果管道一直没有关闭，不会一直阻塞而deadlock，
			//会自动到下一个case匹配, 两个管道取数据的顺序不能保证！
			fmt.Printf("从intChan读取到数据%d\n", v)
			time.Sleep(time.Millisecond * 100)
		case v := <-strChan:
			fmt.Printf("从intChan读取到数据%s\n", v)
			time.Sleep(time.Millisecond * 100)
		default:
			fmt.Printf("都取不到，不玩了, 可以加入自己的逻辑\n")
			time.Sleep(time.Millisecond * 100)
			//break//这里break在select里面不起作用
			//return //太猛了
			//goto lable //只能使用不提倡的goto
			break lable //还是使用break加标签，表示退出这一段for
		}
	}

//lable:
	fmt.Println("主线程正常退出")

}

//go run main.go
//go build -o test.exe
