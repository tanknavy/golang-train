package main
import ("fmt")

func main(){
	a := 23
	b := a //基本类型不是引用类型，b等于a的值
	fmt.Println(a,b) //相同
	a =27 //改变a的值，b会改变吗？
	fmt.Println(a,b) //不相同，因为b基本类型被赋值后有自己的
	fmt.Println(&a,&b) //a,b各自的地址，保存自己的变量

	fmt.Println("----------------")
	var c int = 23
	var d *int = &c //*int声明一个指向int类型指针，保存是地址
	fmt.Println(&c, d, &c==d) //指针变量
	//d, *d, &d :指针变量保存的值(指向的地址)，指针变量指向地址的值，指针自己的地址
	fmt.Println(c, &c, d, *d, &d) //d输出c的16进制地址0xc04200e100
	
	c = 33
	fmt.Println(&c, d)
	fmt.Println(c, &c, d, *d, &d) //d输出c的16进制地址0xc04200e100
	
	*d = 55 //如果指针d指向的值改变呢？c同样会改变
	fmt.Println(c, &c, d, *d, &d) //d输出c的16进制地址0xc04200e100

	aa := [3]int{1, 2, 3} //int默认64位,
	a1 := &aa[0] //0xc0420480c0
	a2 := &aa[1] //0xc0420480c8
	a3 := &aa[2] //0xc0420480d0
	fmt.Printf("%v %p %p %p\n", aa, a1, a2, a3) //a1,a2两个地址相隔0x8,也就是

	//a3 = a3 - 8 //mismatched types *int and int
	fmt.Printf("%v %p\n",a3,a3)
	//unsafe关于地址操作

	var ms *myStruct //指向struct的指针
	fmt.Println(ms, &ms) //nil空值,这是指针没有指向空
	//ms = &myStruct{foo: 43}//&{43}
	ms = new(myStruct) //&{0}，0指foo是int默认值
	fmt.Println(ms,*ms, &ms)
	(*ms).foo = 33 //使用()
	fmt.Println(ms,*ms, &ms)

	
}

type myStruct struct {
		foo int
}