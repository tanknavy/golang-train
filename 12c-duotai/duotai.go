package main

import (
	"fmt"
)

//go run main.go
//go build -o test.exe
//面向对象编程之--多态(poly)，golang里面接口体现多态两种方式: 1)多态参数，2)多态数组

func main() {
	fmt.Println("--多态参数--")
	computer := Computer{}
	phone := Phone{}   //结构体实现了Usb接口
	camera := Camera{} //结构体实现了Usb接口

	computer.Work(phone)  //Usb接口类型变量，不同对象不同动作
	computer.Work(camera) //Usb接口类型变量，不同对象不同动作

	fmt.Println("--多态数组--")
	var usbArr [3]Usb                        //存放实现Usb接口的结构体变量的数组
	usbArr[0] = Phone{name: "iphone 12 pro"} //不同类型放到多态数组里面
	usbArr[1] = Phone{name: "samsung s10"}
	usbArr[2] = Camera{name: "Nikon camera"}
	//如果某个struct有特别方法，类型断言，variable.(Struct)判断variable是否是指向Struct类型的变量，如果是就转换

	fmt.Println(usbArr)

}

type Usb interface {
	//声明两个未实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//Phone实现了Usb接口两个方法，就是实现了这个接口
func (p Phone) Start() {
	fmt.Println("phone start---")
}

func (p Phone) Stop() {
	fmt.Println("phone stop----")
}

func (p Phone) Call() {
	fmt.Println("phone can call----")
}

type Camera struct {
	name string
}

//Camera实现了Usb接口两个方法，就是实现了这个接口
func (c Camera) Start() {
	fmt.Println("camera start---")
}

func (c Camera) Stop() {
	fmt.Println("camera stop----")
}

type Computer struct {
}

//1)golang里面接口体现多态，多态参数
func (c Computer) Work(usb Usb) { //Usb是个接口变量,多态参数，体现出多态的特点
	//通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

//-------------------------------------
//2)golang里面接口体现多态，多态数组
