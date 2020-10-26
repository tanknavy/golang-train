package main

import (
	"reflect"
	"fmt"
)

//reflect遍历struc的字段值，获取结构体的tag标签值，遍历方法，并调用其方法，
//reflect修改struct对象的属性值，refect创建新对象
//reflect.ValueOf()
func main() {
	fmt.Println("-------------------------")
	var a Student = Student{
		Name: "Tom",
		Age: 19,
		Score: 98,
	}
	TestStruct(a)
	fmt.Println("&a地址：",a) //struct本来就是引用
	TestStruct_set(&a) //想修改字段值，需要&
	//TestStruct_new(a) //可传对象也可传对象的指针，
	TestStruct_new(&a)//多了一层Elme()取值
}

type Student struct{
	Name string `json:"name"` //tag，
	Age int `json:"student_age"`
	Score float32
	Sex string
}

func(s Student) Print(){
	fmt.Println("------start-------")
	fmt.Println(s)
	fmt.Println("------end-------")
}

func(s Student) GetSum(n1,n2 int) int{
	return n1+n2
}

//赋值
func(s Student) Set(name string, age int, score float32, sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}){
	//1.获取reflect的Type,Value,Kind
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kin := val.Kind() //注意Type和Kind，Type比如main包下的Student结构体，Kind就是struc
	fmt.Printf("Type:%v, Value:%v, Kind:%v\n",typ,val,kin)
	
	if kin != reflect.Struct { //Kind类别，
		fmt.Println("expect struct")
		return 
	}

	//2.结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n",num)

	for i := 0; i < num; i++ {
		field := val.Field(i) //Value拿到字段
		fmt.Printf("field %d: 值为=%v, 类型=%v\n",i,field,val.Field(i).Kind()) //第几个字段，字段值

		tagVal := typ.Field(i).Tag.Get("json") //Type拿到字段的tag，
		if tagVal != "" { //如果有tag
			fmt.Printf("field %d: tag为=%v\n", i, tagVal)
		}
	}
	//2.1结构体字段赋值或者修改: 反射修改值，第一传入要使用地址&,第二修改时使用Elem()
	//见下面的方法 func(s *Student)
	//val2 := reflect.ValueOf(a)
	//val2.Elem().Field(0).SetString("jack")
	//val2.Elem().FieldByName("Name").SetString("jack")
	//val2.Field(0).Elem().SetString("jack")
	


	//3.结构体有多少个method
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n",num)

	for i := 0; i < numOfMethod; i++ {
		method := val.Method(i)
		fmt.Printf("method方法： %s\n", method)
	}

	//4.结构method的调用
	//方法如何排序?按照方法名称(ascii码比较), 反射
	val.Method(1).Call(nil)	//获取对象的第二个方法，然后调用它,空参使用nil

	//调用结构体的带参方法, 参数必须是relfect.Value的数组，返回还是切片类型
	var params []reflect.Value//声明一个切片slice，数组就是[...]int
	params = append(params,reflect.ValueOf(2)) //
	params = append(params,reflect.ValueOf(3))
	
	res := val.Method(0).Call(params)
	fmt.Println("res=",res[0].Int()) //返回切片

	var params2 [2]reflect.Value//声明一个切片slice，数组就是[...]int
	params2[0] = reflect.ValueOf(12) //index
	params2[1] = reflect.ValueOf(13)
	
	res2 := val.Method(0).Call(params)
	fmt.Println("res2=",res2[0].Int()) //返回Value切片，如果不确定类型，还要断言

	//
	val.MethodByName("Print").Call(nil)
	//fmt.Println("res3=",res3) //返回Value切片，如果不确定类型，还要断言
}


//func TestStruct2(a *Student){ //要操作结构体的字段的值
func TestStruct_set(a interface{}){ //要操作结构体的字段的值，传入地址
	fmt.Println("--------------通过反射修改struct的字段值------------------")
	//2.结构体有几个字段
	//2.1结构体字段赋值或者修改: 反射修改值，第一传入要使用地址&,第二修改时使用Elem()
	val2 := reflect.ValueOf(a)
	//val2.FieldByName("Name").SetString("jack")
	val2.Elem().Field(0).SetString("jack")
	val2.Elem().FieldByName("Name").SetString("TomCat")
	//val2.Field(0).Elem().SetString("jack")
	//fmt.Println("反射修改struct字段name后：",(*a).Name) //a是interface{}类型，a.name不能
	fmt.Println("反射修改struct字段name后：",a) //a是interface{}类型，a.name不能

}

//通过反射来创建对象
func TestStruct_new(a interface{}){//传入一个对象或者对象的地址
	fmt.Println("--------------通过反射new对象------------------")
	typ := reflect.TypeOf(a) //拿到类型
	fmt.Printf("type:=%v,kind=%v\n", typ, typ.Kind())

	//newElem := reflect.New(typ) //通过类型创建指定类型的zero值,如果函数传入的是对象
	newElem := reflect.New(typ.Elem()) //通过类型创建指定类型的zero值,如果函数传入的是对象地址
	
	model := newElem.Interface().(*Student) //空接口再可以assert断言转换
	newElem = newElem.Elem() //拿到新对象的地址,可以复制了
	//newElem.Field(0).SetString("nickName")
	newElem.FieldByName("Name").SetString("Carl")
	newElem.FieldByName("Age").SetInt(39)
	newElem.FieldByName("Score").SetFloat(88.5)
	fmt.Println(model,newElem)
	fmt.Println()
}


//go run main.go
//go build -o test.exe
