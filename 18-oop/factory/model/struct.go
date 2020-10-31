package model

import (
	"fmt"
)

//golang里面没有对象的概念，使用struct代替, 去调用传统oop的extend(有继承), 方法重载，构造函数，析构函数，隐藏this
//struct在var声明变量时，空间就已经分配了，值是各个类型的默认值，struct是值类型
//方法在调用的时候会将调用对象当做实参传递进来,如果想引用，可以使使用struct指针
//Student 定义结构体，类似class，定义时无需逗号，新建对象时需要逗号(最后一个属性赋值也需,)

//如果小写，默认其它包中不能用，使用工厂模式
type Person struct { //大写示可以在其它包表中使用
	Name   string //大写示可以在其它包表中使用
	Gender string
	Age    int
}

//小写的struc只能在包model中使用
//通过工厂模式来解决
type student struct { //大写示可以在其它包表中使用
	Name   string //大写示可以在其它包表中使用
	Gender string
	age    int //小写其它包中不能访问，类似java中private，Get方法封装
}

//给私有的stuct绑定一个方法(但是这个student本来就不可以在外面包中使用)，使用函数
// func (s *student) NewStudent(name string, gender string, age int) (stu student) {
// 	stu = student{
// 		Name:   name,
// 		Gender: gender,
// 		age:    age,
// 	}
// 	return //虽然使用了返回参数(同名)，但还是要return一下
// }

//相当于构造函数
//工厂方法，解决student只能在本包model中使用的问题
func CreateStudent(name string, gender string, age int) (stu *student) { //返回struct指针
	stu = &student{ //地址，这个数据相当于放在堆里面共享的
		Name:   name,
		Gender: gender,
		age:    age,
	}
	return//虽然使用了返回参数(同名)，但还是要return一下
}

func(s *student) SetAge(age int){//封装
	if age >0  && age <150 {
		s.age = age
	}else {
		fmt.Println("年龄不对")//给个默认值
	}
}

//sturct的字段小写，不能在其它包中使用, 体现了封装
func (s *student) GetAge() int{//类似java的public get方法
	//return (*s).GetAge
	return s.age //编译器自动转换
}
