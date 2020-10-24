package main

import (
	"fmt"
)

//Student 定义结构体，类似class，定义时无需逗号，新建对象时需要逗号(最后一个属性赋值也需,)
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

type Teacher struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

//func say(student *Student) string { //这种写法表示一个函数
func (student *Student) say() string { //这种写法表类的方法
	infoStr := fmt.Sprintf("student的信息 name=[%v] age=[%v]", student.name, student.age)
	return infoStr
}

//func say2(student *Student, teacher *Teacher) string { //这种写法表类的方法
func say2(student Student, teacher Teacher) string { //这种写法表类的方法
	//func (student *Student, teacher *Teacher) say2() string { //method has multiple receivers,方法有多个接受者
	infoStr := fmt.Sprintf("student的信息 name=[%v] age=[%v]", student.name, student.age)
	infoStr2 := fmt.Sprintf("teacher的信息 name=[%v] age=[%v]", teacher.name, teacher.age)
	return infoStr + infoStr2
}

func main() {
	//创建实例
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  90,
	}
	//创建实例
	var tea = Teacher{
		name:   "bob",
		gender: "male",
		age:    28,
		id:     1000,
		score:  90,
	}

	fmt.Println(stu.say())
	//fmt.Println(say(&stu))

	fmt.Println(say2(stu, tea))

	tea = Teacher{
		name:   "bob2",
		gender: "male",
		age:    38,
		id:     1000,
		score:  90,
	}

}
