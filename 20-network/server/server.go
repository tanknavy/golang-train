package main

import (
	"fmt"
	//"io"
	"net"
)

//协程，在客户端连接后，专门处理每个客户端的请求
func process(conn net.Conn) { //拿到客户和本机的连接，使用routine处理

	defer conn.Close() //延时关闭，不关闭资源导致服务器
	for {              ////循环接受client发送的数据
		//创建新的切片用于读取客户端输入
		buf := make([]byte, 1024) //字节切片
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有Write(发送)，那么协程就阻塞在这里, tcp协议
		//fmt.Printf("服务器在等待客户端%s, 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //读到了n个字节写入buf切片,一定要截断，否则历史的都一起输出
		if err != nil{
		//if err == io.EOF { //client关闭，我也退出
			fmt.Println("服务器的read err=", err)
			return //如果连接出错，本次连接不玩了
		}
		//3.显示client发送的消息内容到server的终端
		fmt.Print(string(buf[:n])) //截取前n个，客户端已经发送有\n换行
	}
}

func main() {
	fmt.Println("-------------服务器开始监听-----------------")
	listen, err := net.Listen("tcp", "0.0.0.0:8000") //4个0表示ipv4,ipv6都可以，在本地监听
	if err != nil {
		fmt.Println("listen err=", err)
		return //建立监听服务器失败，有错误不玩了
	}

	fmt.Printf("listen suc=%v", listen)
	defer listen.Close() //延时关闭listen

	for { //循环等待客户端来连接
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil { //某个连接失败
			fmt.Println("Accept() error=", err)
			//continue
		} else {
			fmt.Printf("服务器Accept() ok=%v, client客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		//准备一个goroutine，为客户端服务
		go process(conn) //每次conn不一样
	}

}

//go run main.go
//go build -o test.exe
