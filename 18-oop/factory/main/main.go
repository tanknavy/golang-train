package main

import (
	"fmt"
	"goTrain/18-oop/factory/model" //引包的时候从src下面开始, 在model文件夹下struct.go中定义了model包和Person结构体
)

func main() {
	var person = model.Person{ //pkg.struct
		Name:   "bob",
		Gender: "male",
		Age:    28,
	}

	//小写的在其它包中不能使用，使用工厂模式解决
	// var stu = model.student{
	// 	Name : "tom",
	// 	Gender: "male",
	// 	Age : 18,
	// }

	//student小写，使用工厂模式解决
	var stu = model.CreateStudent("tom", "female", 19) //返回指针
	stu.SetAge(21) //封装

	fmt.Println(person)
	//fmt.Println(stu.age)//小写不能在外部包中使用，创建公开的get方法
	fmt.Println(stu.GetAge())
	fmt.Println(*stu)
	
}
