package main

import (
	"fmt"
)

//项目：c/s之间传输消息格式/协议，字符串不能满足各种消息类型，使用struct, 类似b/s之间的map或者json，RESTful
// Message{Type string, Data sting}, 消息类型，消息内容, 为了防止丢包，也可以将消息长度放到这个struct里面 Len uint32
//消息发送流程：
// 1.先创建一个Message的结构体
// 2.msg.Type=登录消息类型
// 3.msg.Data=消息内容(序列化后)
// 4.对msg进行序列化
// 5.在网络传输中，怎么解决丢包？
// 5.1.先给服务器发送msg的长度(多少个字节), 也可以将消息长度放到这个struct里面 Len uint32
// 5.2.再发送消息

//消息接收流程：网络中传输字节
// 1.先接收客户端发送的长度
// 2.根据接收到的长度len，再接收消息本身
// 3.接收时要判断实际接收到的消息内容长度是否等于len
// 4.如果不相等，就有纠错协议
// 5.先反序列化成Message
// 6.再取出meg.Data(string),反序列化Data
// 7.取出loginMsg.userId和loginMsg.userPwd
// 8.到db中比对
// 9.根据比较结果，返回Message{Type string, Data string}, LoginResMesg {code int, error string}

//全局变量，一个表示用户id，一个表示用户密码
var (
	userId  int
	userPwd string
)

//go build -o client.exe goTrain\20-network\oicq\client
func main() {
	//接受用户的选择
	var key int
	//判断是否继续显示菜单
	var loop = true
	for loop { //类似while true,循环等待用户输入
		fmt.Println("--------------欢迎登陆聊天室----------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册新用户")
		fmt.Println("\t\t\t 3 退出聊天室")
		fmt.Println("\t\t\t 4 请选择(1-3)")

		fmt.Scanf("%d\n", &key) //等待输入，从标准输入中读入数字到变量
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册新用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入") //继续循环
		}
	}

	//根据用户输入，继续显示比如二级菜单界面
	if key == 1 {
		//用户要登录
		fmt.Println("请输入用户id:")
		fmt.Scanf("%d\n", &userId) //格式化读入，没有\n的话一串会全没了，这里有坑
		//_, err := fmt.Scanf("%d\n", &userId) //格式化读入，没有\n的话一串会全没了，这里有坑
		//if err != nil {
		//	fmt.Println("输入用户id错误(必须为数字):",err)
		//}
		fmt.Println("请输入用户密码:")
		fmt.Scanln(&userPwd) //要么使用Scaln读一行

		//登录函数，写到另外的文件，比如login.go,都在同一个package下，即使小写也可以直接调用
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {
		fmt.Println("注册新用户")
	}

}
