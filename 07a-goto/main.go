package main

import (
	"fmt"
)

func main()  {

	fmt.Println("statment 01")
	goto label
	fmt.Println("statment 02")
		label:
	fmt.Println("statment 03")
	t01()
	t02()
	fmt.Println("statment 04")
}

func t01(){
	fmt.Println("statment t1")
	fmt.Println("statment t2")
	goto aa
	fmt.Println("statment t3")
	aa:
	fmt.Println("statment 04")
}

func t02(){
	fmt.Println("statment t11")
	goto bb
	fmt.Println("statment t12")
	bb:
	fmt.Println("statment t13")
	fmt.Println("statment t14")
}