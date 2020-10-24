package main

import (
	"fmt"
)

//go run main.go
//go build -o test.exe

//类型断言，variable.(Struct)判断variable指向的是否是Struct类型的变量，如果是就转换
//
func main() {
	//空接口：空方法的接口，因为一切struct都算是绑定了空方法，就相当于实现了interface{}, 有点类似java的Object类？可以引用一切
	var a interface{}             //空接口，一般类型，有点像一切类型的父类
	var point Point = Point{1, 2} //按顺序赋值

	a = point //ok，将struct赋给一个空接口类型
	fmt.Println(a)
	var b Point
	//b = a //直接这样error(相当于父类型赋给子类型),需要断言assert
	//assert类型断言，判断variable指向是否是Struct类型的变量，如果是就相当于将a尝试转为Point类型再赋给b
	b = a.(Point) //要确保原先空接口执行的就是断言的类型，从float32到float64都不行
	fmt.Println(b)

	//带检测的类型断言
	//如果在进行短延时，带上检测机制，不要报panic
	var c interface{}
	var d float32 = 3.14
	c = d
	//一定要c原本指向的类型和要断言的类型一样
	e, ok := c.(float64) //flag成功就true,失败就false
	if ok {
		fmt.Printf("e的类型是%T,转换成功,结果是%v", e, e) //%T查看变量类型，%v是变量的值
	} else {
		fmt.Println("转换失败!")
	}
	
	fmt.Println("转换失败继续执行,不会panic")

}

type Point struct {
	x int
	y int
}
