package main

import (
	"fmt"
	"log"
	"reflect"
)

const (
	_ = iota //iota类似一个range,每次const声明重置，赋值一次+1
)

// 变量,空接口,reflect.Value是可以相互转换的；
// stu := Student{...}
// rv := reflect.ValueOf(stu)
// iv := rv.Interface()
//  v := iv.(Student)

//反射 是在运行时
func reflectTest01(obj interface{}) { //Type,Value, Kind
	//通过反射获取的输入的变量的type(类型), kind(类别)，value
	//1.先获取到reflect.Type接口
	rType := reflect.TypeOf(obj) //反射拿到类型
	fmt.Printf("TypeOf: rType=%v, type=%T\n", rType, rType)

	//2.获取到relfect.Value结构体
	rValue := reflect.ValueOf(obj) //
	fmt.Printf("ValueOf: rValue=%v, type=%T\n", rValue, rValue)

	//3.获取变量对应的Kind
	//Type是类型，Kind是类别(具体分类列表)，它们可能相同(基本类型时)，也可能不同(main.Student, struct)
	rKind1 := rType.Kind()
	rKind2 := rValue.Kind()
	fmt.Printf("rKind=%v, type=%T\n", rKind1, rKind1)
	fmt.Printf("rKind=%v, type=%T\n", rKind2, rKind2)

	//3.返回interface{}引用，必要时.(type)断言转换
	rInter := rValue.Interface()
	fmt.Printf("rInter=%v, type=%T\n", rInter, rInter)

	//将iterface{}通过断言转成需要的类型
	switch rInter.(type) { //type是关键字，类型断言的最佳实践
	case Student:
		fmt.Println("Student类型")
	case int:
		fmt.Println("int类型")
	default:
		fmt.Println("不知道啥类型")
	}
}

func reflectTest02(obj interface{}) { //通过反射修改值，传入必须是地址类型（其它还有值类型）
	//通过反射获取的输入的变量的type(类型), kind(类别)，value，Value.SetXXX改变值

	//2.获取到relfect.Value结构体
	rValue := reflect.ValueOf(obj) //
	fmt.Printf("rValue=%v, type=%T, kind=%v\n", rValue, rValue, rValue.Kind())

	//SetXXX方法, 反射改变变量的值，要使用Elem()
	//调用这个函数时要&地址类型，否则改变不了值，但是rValue变成了指针类型，就不能SetInt()了
	//rValue.SetInt(19) //error, unaddressed,
	//Elem()作用类似 *(&num), 指向变量num的指针，然后指针取值就还是num
	rValue.Elem().SetInt(19) //Elem()返回变量v持有借口保管的值的Value封装，或者持有指针指向值的Value封装

}

func main() {
	fmt.Println("---------------")
	var v int = 3
	var n int = 9
	rv := reflect.ValueOf(v)
	fmt.Printf("rv=%v,type=%T\n", rv, rv) //这个变量打印出来是int 3,但是取不能和int 9相加，因为它是Value类型
	fmt.Println(rv.Int() + int64(n))      //做转换后才能运算
	iv := rv.Interface()
	//fmt.Println(iv + n)//这个变量打印出来是int 3,但是取不能和int 9相加，因为它是interface类型
	fmt.Printf("iv=%v,type=%T\n", iv, iv) //这个变量打印出来是int 3,但是取不能和int 9相加，因为它是interface类型
	vv := iv.(int)                        //assert断言
	fmt.Println(vv + n)                   //断言后才是真正的int，可以和其它int运算

	//基本数据类型
	var num int = 100
	reflectTest01(num)

	//结构体
	stu := Student{name: "Tom", age: 28}
	reflectTest01(stu)
	stu.getIfno()

	//通过反射修改值，注意使用地址&,或者Value.SetXXX
	var num2 int = 9
	reflectTest02(&num2) //传递指正才能真正改变, 但是传入指针后又不能用SetInt()，reflect.Value().Elem().SetInt()
	fmt.Println("num2:", num2)

	//常量
	const (
		_    = iota       //0
		a                 //1
		b, c = iota, iota //2,2
		d    = iota       //3
	)
	fmt.Println(a, b, c, d)

	//日志
	log.Println("log...")

}

type Student struct {
	name string
	age  int
}

func (stu Student) getIfno() {
	fmt.Printf("name=%s,age=%d\n", stu.name, stu.age)
}

//go run main.go
//go build -o test.exe
