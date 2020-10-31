package main

import (
	"fmt"
)

//interface降低耦合，接口是指针(引用类型)，struct是值传递
//可以把任何一个变量赋给空接口interface{}，类似java中的Object?
//golang中通过interface体现多态
type Usb interface { //接口相当于一个规范，松耦合，高内聚
	Start() //接口中未实现的方法，可以有输入参数和返回值
	Stop()
}

type Usb2 interface { //实现接口没有implement关键字，只是方法名
	Start()
	Stop()
	//Test()
}

type Phone struct {
	Name string
}

//struct实现interface，实现了Usb的全部方法, Usb2中也有同名方法，它也实现了Usb2这个接口吗？是的
func (p Phone) Start() {
	fmt.Println("Phone begin to work...")
}
func (p Phone) Stop() {
	fmt.Println("Phone stop to work...")
}

type Camera struct {
	Name string
}

//struct实现interface，实现了Usb的全部方
func (c Camera) Start() {
	fmt.Println("Camera begin to work...")
}
func (c Camera) Stop() {
	fmt.Println("Camera stop to work...")
}

type Computer struct {
}

//只要是实现了Usb接口(实现了接口声明所有方法),接口是引用类型，不同于sturct的值传递
//参数类型不同，体现不同状态
func (c Computer) Work(usb Usb2) { //接受一个interface类型变量,根据具体的类型调用相应的方法，多态！
	usb.Start() //通过接口调用方法
	usb.Stop()
	fmt.Printf("usb 接口引用=%v\n", usb)

}

type AInterface interface { //接口
	SayHi()
	//name string //golang接口中不能有变量，java中interface里面变量都是常量
}
type integer int

//除了struct,自定义数据类型也可以实现接口
func (i integer) SayHi() { //integer这个自定义类型也实现了A接口
	fmt.Println("integer实现了A接口")
}

type BInterface interface { //接口
	SayHey()
}

type Monster struct {
}

func (m Monster) SayHi() { //
	fmt.Println("A interface method")
}

func (m Monster) SayHey() {
	fmt.Println("B interface method")
}

type CInterface interface { //接可以继承接口，如果这些接口有共同方法怎么办？type可以实现不同接口的相同方法,但是不能继承有公共方法的多个接口
	AInterface //接口继承其它接口
	BInterface //接口继承其它接口
	SayHello() //自己的方法
}

type Animal struct {
}

func (a Animal) SayHi() { //
	fmt.Println("A interface method")
}

func (a Animal) SayHey() {
	fmt.Println("B interface method")
}

func (a Animal) SayHello() {
	fmt.Println("C interface method")
}

type T interface{} //定义T为空接口类型
type S struct{}

func main() {
	//1.实现接口，多态引用
	var phone = Phone{"iphone 12 pro"}
	var camera = Camera{"canon ex8"}
	var computer = Computer{}

	computer.Work(phone) //把结构体变量传给了Usb
	fmt.Printf("phone=%p\n", &phone)
	computer.Work(camera)
	fmt.Printf("camera=%p\n", &camera)

	//2.除了struct, 自定义类型也可以实现接口
	var i integer = 10
	i.SayHi()
	var a AInterface = i //接口类型定义变量
	a.SayHi()

	//3.实现多个接口
	var monster AInterface = Monster{}
	monster.SayHi()
	//monster.SayHey() //错误,只能调用对应接口下的方法
	var monster2 BInterface = Monster{}
	monster2.SayHey()

	//4.接口继承接口, 必须实现所有方法
	var animal CInterface = Animal{}
	animal.SayHello()
	var animal2 AInterface = Animal{}
	//animal2.SayHello() //错误，只能调用对应接口下的方法
	animal2.SayHi()

	//5.接口是指针(引用类型)，如果没有对interface初始化就使用会输出nil
	var animal3 CInterface
	fmt.Println(animal3) //只定义没有初始化，输出为nil

	//6.空接口interface{}没有任何方法，所以所有类型都实现了空接口！interface{}可以引用全部变量
	// 可以将任何类型的变量赋给空接口类型
	var t T = phone //T是interface{},这样可以吗？yes，空接口可以接受任何数据类型！
	fmt.Println(t)
	//t.Start() //错误，t是空接口类型没有任何方法
	var t2 interface{} = 28 //随便赋值
	t2 = "Millions"         //可以随便赋值
	var t3 float32 = 3.14
	//t3 = "hello" //错误,类型不匹配
	var t4 S //空struct,体会下区别

	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

}
