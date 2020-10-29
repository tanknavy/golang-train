package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"goTrain/20-network/oicq/common/message"
	"io"
	"net"
)

//server读取client的数据包
func readPkg(conn net.Conn) (msg message.Message, err error) { //golang里面返回值会自动产生默认的，程序中可以直接用
	buf := make([]byte, 8096) //缓存slice用于读取客户端数据
	fmt.Println("从客户端读取到的数据...")
	//1.读消息头，判断下次实际要读取的消费在buf中的位置
	//Read如果读不到东西会阻塞(所以在channel中可以循环读)，前提是有效连接，前面是4字节的数据包长度内容
	_, err = conn.Read(buf[:4])
	//if n != 4 || err != nil {    //n实际读取了几个字节
	if err != nil { //n实际读取了几个字节
		//fmt.Println("conn.Read err=", err) //conn断开时也会err
		//err = errors.New("read pkg head error") //自定义error,这个err和返回值形参同名，就会返回它，注意有的类型要使用&赋值给返回的形参
		return
	}
	//fmt.Println("读到的buf=", buf[:4])

	//2.根据读到的buf[:4]，转成一个uint32类型，读取buf里面相应区间的数据
	var pgkLen uint32 = binary.BigEndian.Uint32(buf[:4]) //将一个4字节的表示长度的数字转成unit32，比如[0 0 1 0]转成256
	//n, err := conn.Read(buf[:pgkLen+4]) //从conn套接字中读pkgLen个字节放到buf中，不是从buf从读取！！！
	n, err := conn.Read(buf[:pgkLen])   //从conn套接字中读pkgLen个字节放到buf中，不是从buf从读取！！！
	if n != int(pgkLen) || err != nil { //n实际读取了几个字节
		//fmt.Println("conn.Read err=", err) //conn断开时也会err
		//err = errors.New("read pkg body error") //自定义error
		return //直接返回err(同返回值新参同名)
	}
	//fmt.Println("读到的buf=", buf[:pgkLen]) //字节数组

	//3.反序列化成Message类型
	err = json.Unmarshal(buf[:pgkLen], &msg) //反序列化成Message类型，直接写到返回的msg变量，不加&会导致返回的msg是个空的
	if err != nil {                          //n实际读取了几个字节
		//fmt.Println("json.Unmarshal err=", err) //conn断开时也会err
		err = errors.New("message deser error") //自定义error
		return
	}

	return //golang里面返回值会自动产生默认值，程序里面可以直接用，一定要return,它会返回程序中同名的变量
}


//处理和客户端通讯
func process(conn net.Conn) {
	//每个客户端延时关闭，否则奇怪的问题
	defer conn.Close()

	//循环读取客户端发送的信息,长连接
	for {
		//这里将读取数据包，直接封装成一个函数readPkg(), 输入conn,返回msg,err
		msg, err := readPkg(conn)
		if err != nil { //n实际读取了几个字节
			if err == io.EOF { //client关闭的conn错误
				fmt.Println("client退出了，server端连接也退出..")
			} else { //别的错误
				fmt.Println("readPkg(conn) err=", err) //conn断开时也会err
			}
			return //读取客户端错误，不玩了
		}

		fmt.Println("msg=", msg)

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
		conn, err := listen.Accept() //循环处理每个client的连接请求
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功，启动一个routine和客户端保持通讯...
		go process(conn) //协程

	}

}
