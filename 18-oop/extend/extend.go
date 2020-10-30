package main

import (
	"fmt"
)

//实现接口 vs 继承
//接口实现可以在不破坏继承关系的基础上对结构体进行扩展功能，是一种解耦的方式
//可以认为实现接口是对继承机制的补充，结构比继承更加灵活
//inherit：价值在于解决代码的复用性和可维护性，满足is-a的关系
//interface：设计,设计好各种规范(方法),让其他自定义类型是实现这些方法，满足like -a的关系
//多态：变量(实例)具有多种形态，在golang中多态是通过接口实现的，可以按照统一的接口来调用不同的实现

//任何类型都实现了空接口,因为空接口什么方法都没有,所以interface{}可以引用一切类型，
//variable.(type)使用assert类型断言将interface{}类型变量转成本来的类型

func main() {
	monkey := SuperMonkey{
		//Monkey{name: "wukong",}, //只有父类属性时，匿名
		Monkey: Monkey{name: "wukong"}, //混合父类字段赋值，和自己属性赋值时，要使用:
		title:  "齐天大圣",
	}
	monkey.climb() //继承来的功能
	monkey.fly()   //
	monkey.swim()  //
}

type Monkey struct { //结构体，父类
	name string
}

//给Monkey这个结构体绑定一个方法
func (monkey *Monkey) climb() {
	fmt.Println(monkey.name + "猴子生来会爬树(子类继承父类)")
}

type SuperMonkey struct { //结构体，子类
	Monkey //匿名继承,父类的属性和方法就可以使用
	title  string
}

//想扩展功能又不想破坏继承关系
//使用interface扩展功能
type Bird interface {
	fly()
}

type Fish interface {
	swim()
}

type shenxian interface {
	power()
	magic()
}

//实现接口,就是实现接口里面声明的方法
func (monkey *SuperMonkey) fly() {
	fmt.Println(monkey.name + "学会筋斗云(实现接口)")
}

//实现接口,就是实现接口里面声明的方法
func (monkey *SuperMonkey) swim() {
	fmt.Println(monkey.name + "学会潜游海底(实现接口)")
}

func (monkey *SuperMonkey) power() {
	fmt.Println(monkey.name + "学会神仙术的法力(实现接口)")
}

func (monkey *SuperMonkey) magic() {
	fmt.Println(monkey.name + "学会神仙术的魔法(实现接口)")
}
