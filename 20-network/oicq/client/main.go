package main

import (
	"fmt"
)

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
