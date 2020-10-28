package main

import (
	"fmt"
	"net"
)

//处理和客户端通讯
func process(conn net.Conn){
	//读取客户端发送的信息
	
}

func main(){
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net.Listen.err=",err)
	}

	//一旦监听成功，等待客户端连接服务器
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("listen.Accept err=",err)
		}

		//一旦连接成功，启动一个routine和客户端保持通讯...
		go process(conn)

	}
}