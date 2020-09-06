package main //把test.go文件归属到main包,go语言的每个文件都要属于一个包，
import ("fmt"; "unsafe"; "app01/variables") //引入一个包fmt

//指针:p(表示保存的变量-即地址),*p(指向的变量),&p(自己的地址)
func testPtr(num *int){ //*号表示指针类型
	*num = 20 //指针指向的值
	fmt.Println(num)
	
	
}

//主函数
func main() {
	fmt.Println("hello Golang")
	var p *int //p是一个指针指向一个整数型，指针保存的是变量的地址

	i := 999
	j := 80
	p = &i
	testPtr(p) //调用函数
	fmt.Println("----------------")
	fmt.Println(*p) //p指针指向的值
	fmt.Println(p) //p指针指向的内存地址
	fmt.Println(&p) //p指针自己的内存地址
	fmt.Println("----------------")
	p = &j
	fmt.Println(*p) //p指针指向的值
	fmt.Println(p) //p指针指向的内存地址
	fmt.Println(&p) //p指针自己的内存地址

	fmt.Println("----------------")

	fmt.Println(**&p) //p指针自己的地址->这个地址(指针)保存的值(另外一个地址)->另外地址保存的值(一个整数)
	fmt.Println(0xc042004030) //结果将16进制转成了10进制
	
	fmt.Println("----------------")
	var ptr uintptr = 0xc042004030 //地址类型变量
	q := unsafe.Pointer(ptr) //
	fmt.Println(&q) //q指针自己的地址
	fmt.Println(q) //q指向的地址
	//fmt.Println(*q) //q指向的地址所保存的值
	
	fmt.Println("----------------")
	variables.varTest()


}