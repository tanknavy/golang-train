package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goTrain/20-network/oicq/common/message" //GOPATH的src下全路径的
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	// //开始定协议
	// fmt.Printf("userId=%d,userPwd=%s\n",userId, userPwd)
	// return nil

	//1.连接到server
	//应该读配置文件,比如gonfig,https://medium.com/@onexlab.io/golang-config-file-best-practise-d27d6a97a65a
	conn, err := net.Dial("tcp", "localhost:8889") //
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return //连不上就玩不了
	}
	//延时关闭链接,如果没写，后面有问题还不好调试
	defer conn.Close()

	//2.准备通过conn发送消息(结构体)给server
	var msg message.Message
	msg.Type = message.LoginMsgType

	//3.创建一个LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	//4.msg.Data = loginMsg //error,需要先序列化成json串
	data, err := json.Marshal(loginMsg) //返回byte
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.data赋给msg.Data
	msg.Data = string(data) //[]byete要转string

	//6.将msg进行序列化
	data, err = json.Marshal(msg) //相同类型，不用:=，返回的[]byte
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.这时data就是可以发送到server的消息，记得避免丢包，要先发送消息长度
	//7.1 将data的长度发送给服务器
	// 先获取data长度->转成一个表示长度的byte切片，
	//encoding/binary实现数字与字节序列的转换
	var pkgLen uint32 = uint32(len(data)) //data是字节数组，数据长度无符号32位,强转
	var buf [4]byte //准备一个数组，为啥长度4，因为32bit就是4byte
	//var buf []byte = make([]buf, 4, 4)
	//var buf []byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //将字节的长度数字转为字节slice,比如长度8转成32位的二进制，每8bit(一byte)为slice中一个元素
	//发送长度, 可以看到slice从arr中切片，slice中元素是指针，slice的底层array，
	n, err := conn.Write(buf[:4]) //需要发送[]byte，返回多少字节被发送，err
	if n != 4 || err != nil {   //如果发送的byte长度不为4或者有错误
		fmt.Println("conn.Write(buf) fail:", err)
		return //
	}

	fmt.Printf("客户端 发送消息的长度=%d,内容=%s",len(data),string(data)) //tcp是一个长连接
	return //到这里err一直为nil,成功返回
}
