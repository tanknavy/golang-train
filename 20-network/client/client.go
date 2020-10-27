package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("------------------------------")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn 成功=", conn)

	//从标准终端读入
	reader := bufio.NewReader(os.Stdin)

	//如果要连续输入，使用for循环直到exit
	//从终端读取一行用户输入，准备发给送服务器
	for {
		line, err := reader.ReadString('\n') //读入字符串并加上\n
		if err != nil {
			fmt.Println("readString err=", err)
		}

		line = strings.Trim(line, " \r\n") //为了检测是否输入了exit
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		//将line发送给服务器
		_, err = conn.Write([]byte(line + "\n")) //字符串强转字节数组，写给server

		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		//fmt.Printf("客户端发送了%d字节数据并退出",n)
	}

}

//go run main.go
//go build -o test.exe
