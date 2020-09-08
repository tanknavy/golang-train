package main 

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Stacy"
	sayMessage(greeting, &name)//指针
	fmt.Println(name) //使用指针后，name被改变了

	tmp := sum("The sum is:",1,2,3,4,5)
	fmt.Printf("%v %T", tmp, tmp)

	sum2("func w/o return:",1,2,3,4,5)

	func(){ //匿名函数
		fmt.Println("I'm 匿名函数")
	}() //()invoke function

	//函数作为变量, 无需形参
	var f func(float64, float64) (float64,error)//error也是一种类型 
	f = func(a,b float64) (float64,error) { 
		fmt.Println("函数变量")
		if b == 0.0 {
		//panic("cannot provide zero as second value")
		return 0.0, fmt.Errorf("cannot provide zero as second value")
	}
	return a / b, nil //返回值和nil错误
	}
	f(5.0,3.0)


	g := greeter { //实例化一个struct
		greeting: "hello,struct with method",
		name: "go",
	}
	g.greet()//调用struct的method


	g.greetUpdate("great","alex")
	fmt.Println("struct g new name:",g.name)

	type counter int

	d,err := div(5.0, 0.0)//带有错误返回的function
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

//------------------------------------------------------
//使用用指针变量后，
func sayMessage(greeting string, name *string){
	fmt.Println(greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(msg string, values ...int) *int{ //可变参数作为slice传入
	fmt.Println(values)
	res := 0
	for _,v := range values{
		res += v
	}
	fmt.Println(msg, res)
	return &res //函数返回地址
}

func sum2(msg string, values ...int) (res int){ //输出参数可以使用
	fmt.Println(values)
	for _,v := range values{
		res += v //输出形参中res默认0
	}
	fmt.Println(msg, res)
	return
}


func div(a,b float64) (float64, error){//返回类型可推断
	if b == 0.0 {
		//panic("cannot provide zero as second value")
		return 0.0, fmt.Errorf("cannot provide zero as second value")
	}
	return a / b, nil //返回值和nil错误
}

type counter int //method也可以和除了struct的其它类型搭配
//struct + func(s struct)模拟了java中的class
type greeter struct {//定义greeter为结构体
	greeting string
	name string
}

func (g greeter) greet(){ //结构体方法，现在不叫function而叫method
	fmt.Println(g.greeting,g.name)
	//struct作为变量传入是copy方式，这里改变不会影响原struct
	//除非将func (g *greeter)定义指针方式 
	g.name = "new Name" 
}

func (g *greeter) greetUpdate(greeting string, name string){ 
	//fmt.Println(g.greeting,g.name)
	g.greeting = greeting
	g.name = name
}