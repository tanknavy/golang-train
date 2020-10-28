package main

import (
	"fmt"
	"net"
)

//处理和客户端通讯
func process(conn net.Conn) {
	//每个客户端延时关闭，否则奇怪的问题
	defer conn.Close()

	//循环读取客户端发送的信息,长连接
	for {
		buf := make([]byte, 8096) //缓存slice用于读取客户端数据
		fmt.Println("读取客户端读取到的数据...")
		n, err := conn.Read(buf[:4]) //注意Read如果读不到东西会阻塞(所以在channel中可以循环读)，前面是4字节的数据包长度内容
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=", err) //conn断开时也会err
			return
		}
		fmt.Println("读到的buf=", buf[:4])
	}

}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen.err=", err)
	}

	//一旦监听成功，等待客户端连接服务器
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功，启动一个routine和客户端保持通讯...
		go process(conn)

	}

}
