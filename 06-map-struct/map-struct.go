package main

import (
	"fmt"
	"reflect"
)

//首字母大写表示global/public, 小写表示private
//struct如何系列化？json.Marshal(&struct)返回byte的slice
func main() {
	populations := map[string]int{ //定义map[string]int
		"ca": 29, //格式类似json
		"tx": 22,
		"wa": 18, //逗号必须
	}
	m := map[[3]int]string{}

	states := make(map[string]int, 10) //make定义一个map

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
	type Doctor struct { //定义结构体(类似class)，无需任何标点符号
		id         int
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
		Animal: Animal{name: "emu", origin: "Australia"}, //赋值方式一
		speed:  48,
		canFly: true,
	}
	//b.name = "emu" //可以先创建Bird{}再逐个字段赋值
	//b.origin = "Australia"
	fmt.Println(b)

	t := reflect.TypeOf(Animal{})     //反射struc
	field, _ := t.FieldByName("name") //struc字段
	fmt.Println(field.Tag)            //字段的tag

}
