//package model2 //同一个文件夹下不同两个包名！！
package model

import (
	_ "fmt"
)

//golang里面没有对象的概念，使用struct代替, 去调用传统oop的extend(有继承), 方法重载，构造函数，析构函数，隐藏this
//struct在var声明变量时，空间就已经分配了，值是各个类型的默认值，struct是值类型
//方法在调用的时候会将调用对象当做实参传递进来,如果想引用，可以使使用struct指针
//Student 定义结构体，类似class，定义时无需逗号，新建对象时需要逗号(最后一个属性赋值也需,)

//如果小写，默认其它包中不能用，使用工厂模式
type Person2 struct { //大写示可以在其它包表中使用
	Name   string //大写示可以在其它包表中使用
	Gender string
	Age    int
}

type student2 struct { //大写示可以在其它包表中使用
	name   string //大写示可以在其它包表中使用
	gender string
	age    int
	id     int
	score  float64
}