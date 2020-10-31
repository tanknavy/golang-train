package main

import (
	"fmt"
)

//golang里面没有对象的概念，使用struct代替, 去调用传统oop的extend(有继承), 方法重载，构造函数，析构函数，隐藏this
//struct在var声明变量时，空间就已经分配了，值是各个类型的默认值，struct是值类型
//方法在调用的时候会将调用对象当做实参传递进来,如果想引用，可以使使用struct指针
//Student 定义结构体，类似class，定义时无需逗号，新建对象时需要逗号(最后一个属性赋值也需,)
//结构体嵌套匿名/有名结构体或数据类型
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

//-------------------struct绑定方法--------------
//实现了String方法，那么fmt.Println默认调用这个变量的String输出
func (t *Teacher) String() string { //调用时记得传入指针
	str := fmt.Sprintf("name=[%v] age=[%v] id=[%v] score=[%v]\n", t.name, t.age, t.id, t.score) //Sprintf格式化返回字符串而不是打印
	return str
}

type Circle struct {
	radius float64
}

//struct绑定方法
func (c Circle) area() float64 { //方法回家调用对象c当做实参传递进来
	return 3.14 * c.radius * c.radius //golang底层做了转换
}

func (c *Circle) area2() float64 { //引用
	//return 3.14 * (*c).radius * (*c).radius //本来应该这样写
	//c.radius = 10 //因为是指针，也会修改c对象的radius属性值
	return 3.14 * c.radius * c.radius //编译器底层做了转换，可以这样简写
}

//还可以对非struct绑定方法，这样integer就是int的别名
type integer int //可以对struct绑定方法，也可以对自定义类型绑定，比如int,float32

//func say(student *Student) string { //函数
func (student *Student) say() string { //表示给struct绑定一个方法，传入了结构体指针
	infoStr := fmt.Sprintf("student的信息 name=[%v] age=[%v]", student.name, student.age) //可以将打印内容传递给变量
	return infoStr
}

//func say2(student *Student, teacher *Teacher) string { //函数
func say2(student Student, teacher Teacher) string { //函数
	//func (student *Student, teacher *Teacher) say2() string { //method has multiple receivers,方法有多个接受者
	infoStr := fmt.Sprintf("student的信息 name=[%v] age=[%v]", student.name, student.age)
	infoStr2 := fmt.Sprintf("teacher的信息 name=[%v] age=[%v]", teacher.name, teacher.age)
	return infoStr + infoStr2
}

//-----------------struct的继承--------------
type A struct {
	Name string
	age  int
}

type B struct {
	Name string
	age  int
}
type C struct { //C嵌套了匿名结构体，也就是继承, 继承多个就是多重继承(尽量不要使用多重继承)
	A
	B
	Name string //结构体和匿名结构体有同名属性或方法，编译器采用就近原则，如果匿名结构体之间同名，访问时指定结构体.属性
}

type D struct { //嵌套有名结构体
	a A //嵌套了有名结构体，这种模式就是组合关系,这时调用时必须使用a.属性
}

type E struct { //嵌套结构体指针
	*A   //嵌套结构体指针
	*B   //嵌套结构体指针
	Name string
}

type F struct {
	A   //除了嵌套struct来继承，
	int //嵌套基本数据类型,访问时f.int,不能有第二个，如果一定要有必须取名
	age int
}

//---------

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

	//-------------------struct的继承--------------
	cir := Circle{5.0}
	fmt.Println(cir.area())
	//fmt.Println((&cir).area2()) //标准的写法，但编译器做了优化了，可以简写
	fmt.Println((&cir).area2()) //可以这样简写

	var c C
	c.Name = "tom"
	fmt.Println(c,c.Name, c.A.Name, c.B.Name)//字段同名如何访问，都有同名字段要加struct，如果唯一属性可以直接点

	var d D
	d.a.Name = "jack" //嵌套有名结构体，访问必须加名称.属性,方法依然
	fmt.Println(d)

	//可以在定义变量时指定匿名结构体的字段值
	c2 := C{
		A{"adm", 19},
		B{"bob", 20},
		"carl"}
	fmt.Println(c2)

	e := E{ //嵌套结构体指针
		&A{"adm", 19},
		&B{"bob", 20},
		"carl"}
	fmt.Println(e)
	fmt.Println(e.A, *e.A) //取地址，取值

	f := F{
		A{"adm", 19},
		175, //嵌套基本数据类型
		28, //字段
	}
	f.int = 181//给匿名嵌套
	f.age = 29
	fmt.Println(f.int,f.age) //访问嵌套基本数据类型

}
