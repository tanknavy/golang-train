package main

import (
	"fmt"
)

// const identifier [type] value, 编译器在编译期间就要确定
//常量必须赋值，只能修饰bool, 数值类型，string类型
//iota类似一个range,每次const声明重置，赋值一次+1
//很多语言要求常量命名全大写，golang没有要求，但是依然通过首字母大小控制常量的访问范围
const ( //多个常量使用(),写在外面就是全局，写在fun内面就是局部
	//enumerated successive integer constants 0, 1, 2
	//It resets to 0 whenever the word const appears in the source code
	//and increments after each const specification.
	_ = iota //初始值0,后面从1开始
	a = iota //enumerated successive integer constants 0, 1, 2
	b //继续依次+1
	c //3
	d //继续依次+1,4
	e,f = iota, iota //这样同一行就会相等,5
	g = iota //6
)

func main() {

	const c1 float64 = 3.14
	//const c1 float64 = math.Sin(3.14)  //常量不可再赋值,scala中val关键字
	// iota枚举0,1,2...，在const关键之声明时重置为0， 每次赋值再加一
	fmt.Printf("%v,%T\n", c1, c1)
	fmt.Printf("a:=%v\n", a) //1
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)//5
	fmt.Printf("%v\n", f)//5
	fmt.Printf("%v\n", g)//5
}
