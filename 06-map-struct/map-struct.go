package main

import (
	"fmt"
	"reflect"
)

//首字母大写表示global/public(可以在其它package中使用), 小写表示private(本包中使用)
//map是引用类型，struct是值类型，golang中slice,map,channel是引用类型类型(默认为nil即还没有开辟空间，使用make)，其它是值类型(有相应默认值，可使用new)
//struct如何系列化？json.Marshal(&struct)返回byte的slice
func main() {
	populations := map[string]int{ //定义map[string]int
		"ca": 29, //格式类似json
		"tx": 22,
		"wa": 18, //逗号必须
	}
	m := map[[3]int]string{}

	states := make(map[string]int, 10) //make定义一个map，开辟空间，map可自动增长

	fmt.Printf("%v,%T\n", populations, populations)
	fmt.Println(populations)
	fmt.Printf("%v,%T\n", m, m)

	fmt.Printf("%v,%T\n", states, states)

	populations["ca"] = 33 //取值并重新赋值
	fmt.Println(populations["ca"])
	fmt.Println(populations["ny"]) //不存在的key返回默认值0

	pop, ok := populations["fl"] //ok返回bool是否有值
	fmt.Println(pop, ok)         //0 false

	fmt.Println(len(populations)) //元素个数
	delete(populations, "az")     //删除map中一个元素

	//----------------------------------------------
	//定义结构体(类似class)，无需任何标点符号
	type Doctor struct { //大写示可以在其它包表中使用
		id         int //定义结构体(类似class)，无需任何标点符号
		name       string
		companions []string
	}

	aDoctor := Doctor{ //创建一个Doctor类型的变量
		id:         3, //struct的字段赋值, 格式类似json
		name:       "Joe",
		companions: []string{"Liz", "Jack"}, //最后需要,
	}

	//注:可以不定义type struct定义，直接创建一次性(匿名)的struct对象
	bDoctor := struct{ name string }{name: "Greg"} //匿名struct{字段定义}{字段赋值}, 类似java匿名内部类
	cDoctor := bDoctor
	cDoctor.name = "Bob"

	fmt.Printf("%v,%T\n", aDoctor, aDoctor)
	fmt.Printf("%v,%T\n", bDoctor, bDoctor)
	fmt.Printf("%v,%T\n", cDoctor, cDoctor)

	//结构体可以被嵌套(类似继承,匿名继承)
	type Animal struct {
		name   string `required max:"10"` //使用tag, 还有json:"name"指定json序列化时的属性名 
		Origin string `json:"origin_"` //首字母大小表示export
	}

	type Bird struct {
		Animal //结构体集成另外一个结构体(类似继承,匿名继承)
		speed  float32
		canFly bool
	}

	b := Bird{ //继承的struc不能直接这样写继承的字段
		//name : "emu", //继承后，未知字段
		//origin : "Australia",
		Animal: Animal{name: "emu", Origin: "Australia"}, //赋值方式一
		speed:  48,
		canFly: true,
	}
	//b.name = "emu" //可以先创建Bird{}再逐个字段赋值
	//b.origin = "Australia"
	fmt.Println(b)

	//反射
	t := reflect.TypeOf(Animal{})     //反射struc
	field, _ := t.FieldByName("name") //struc字段
	fmt.Println(field.Tag)            //字段的tag

	//map和struct是值类型还是引用类型?
	fmt.Println("-------map&struct的类型-------------")
	populations2 := populations
	fmt.Println(populations,populations2) //相同，说明map是引用类型
	//fmt.Printf("%p,%p\n",&populations,&populations2) //相同，说明map是引用类型&

	dDoctor := aDoctor
	dDoctor.name = "Carl"
	fmt.Println(aDoctor,dDoctor)//不同，struct是值类型
	//fmt.Printf("%p,%p\n",&aDoctor,&dDoctor)

	chan1 := make(chan int, 3)
	chan2 := chan1
	chan1 <- 1
	fmt.Println(chan1, chan2) //地址相同
	//fmt.Printf("%p,%p\n",&chan1, &chan2)

	//不同结构体变量的字段是独立的，
	fmt.Println("-------struct不同字段类型的默认值,引用类型默认nil-------------")
	type Person struct{//基本数据类型可以使用new, 引用类型要使用make开辟空间
		name string //默认""空串
		age int //int,float默认0
		xx bool //默认false
		score [3]int //指定个数就是array，默认[x x x]
		ptr *int //指针, 默认nil，基本数据类型使用new开辟空间
		hobby []string //slice,默认nil, 输出[], 测试==nil, 使用要么声明时赋值，make开辟空间，或者append
		theMap map[string]string //默认nil,输出map[],使用前make开辟空间
	}
	var person Person
	fmt.Println(person)
}
