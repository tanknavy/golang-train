package main

import (
	"fmt"
	"goTrain/20-network/oicq/common/message"
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	// //开始定协议
	// fmt.Printf("userId=%d,userPwd=%s\n",userId, userPwd)
	// return nil

	//1.连接到server
	conn, err := net.Dial("tcp", "localhost:8889") //应该读配置文件，https://medium.com/@onexlab.io/golang-config-file-best-practise-d27d6a97a65a
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return //连不上就玩不了
	}

	//2.准备通过conn发送消息(结构体)给server
	var msg message.Message
	msg.Type = message.LoginMsgType

	//3.创建一个LoginMsg结构体
	var loginMsg message.LoginMsg

}
