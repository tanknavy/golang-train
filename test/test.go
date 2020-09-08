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

}

func test01() string{
	//Aa := "hello,go"
	return "global scope"
}