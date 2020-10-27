package main
import "fmt"

func main(){
	var a []byte = []byte("hello") //将hello字符串转换为字节slice
	a = append(a,'k')
	fmt.Printf("%v %T\n", a,a)

	var B []byte = []byte{'a','b','c'} //定义个slice,包含3个元素
	B = append(B,'k')
	fmt.Printf("%v %T\n", B,B)

	
	test01()

	person := Person{name:"Jr. Bob", age:29}
	person.show()
	fmt.Printf("%v, %T\n", person, person)

	var pp Generic = Person{name:"Sr. Bob", age:39}//接口做类型引用
	pp.show()

	
}

func test01() string{
	//Aa := "hello,go"
	return "global scope"
}

type Person struct{
	name string
	age int
}

type Generic interface{
	show()
}

func (p Person) show() {
		fmt.Println("my name is ", p.name)
}
