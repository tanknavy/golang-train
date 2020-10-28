package main

import ("fmt";"unsafe")//go编程无需;结尾

var ( //变量和作用域
	a string= "aa is String定义变量" //字符串变量
	b int  = 33 //作用域在main包,package scope
	D int = 95 //大写变量expose到外面，global scope
)

func main() {
	var i int //定义变量整形，作用域在此func scope
	i = 24
	var j float32 =3.14 //定义变量浮点型, float32(i)整形转浮点型
	k:="hello" //定义变量字符串型，类型，赋值，快捷方式
	fmt.Println(a,b)
	fmt.Printf("%v,%T\n",D,D)//打印变量j，变量类型
	fmt.Println(i,j,k)
	fmt.Printf("%v,%T\n",j,j)//打印变量j，变量类型
	fmt.Printf("%v,%T\n",k,k)//打印变量k，变量类型
	
	fmt.Println("-----------------------------")
	var ptr uintptr = 0xc042004030 //地址类型变量
	q := unsafe.Pointer(ptr) //不安全的指针类型
	fmt.Printf("%v,%T\n",q,q)//打印变量和变量类型
	fmt.Println("------------------------------")
	fmt.Printf("%v,%T\n",&i,&i)//打印变量的地址，地址类型

}

