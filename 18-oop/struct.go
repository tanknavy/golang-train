package main

import (
	"fmt"
)

//golang里面没有对象的概念，使用struct代替, 去调用传统oop的extend(有继承), 方法重载，构造函数，析构函数，隐藏this
//struct在var声明变量时，空间就已经分配了，值是各个类型的默认值，struct是值类型
//方法在调用的时候会将调用对象当做实参传递进来,如果想引用，可以使使用struct指针
//Student 定义结构体，类似class，定义时无需逗号，新建对象时需要逗号(最后一个属性赋值也需,)
type Student struct { //大写示可以在其它包表中使用
	name   string //大写示可以在其它包表中使用
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

type Circle struct {
	radius float64
}

func (c Circle) area() float64 { //方法回家调用对象c当做实参传递进来
	return 3.14 * c.radius * c.radius //golang底层做了转换
}

func (c *Circle) area2() float64 {
	//return 3.14 * (*c).radius * (*c).radius //本来应该这样写
	return 3.14 * c.radius * c.radius //golang底层做了转换，可以这样简写
}

type integer int //可以对struct绑定方法，也可以对自定义类型绑定

//func say(student *Student) string { //这种写法表示一个函数
func (student *Student) say() string { //这种写法表类的方法
	infoStr := fmt.Sprintf("student的信息 name=[%v] age=[%v]", student.name, student.age) //可以将打印内容传递给变量
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

	c := Circle{5.0}
	fmt.Println(c.area())
	fmt.Println((&c).area2()) //标准的写法，但底层优化了，可以简写

}
