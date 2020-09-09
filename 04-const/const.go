package main

import (
	"fmt"
)

const (
	//enumerated successive integer constants 0, 1, 2
	//It resets to 0 whenever the word const appears in the source code
	//and increments after each const specification.
	_ = iota //初始值0,后面从1开始
	a = iota //enumerated successive integer constants 0, 1, 2
	b = iota
	c = iota
	d
	e
)

func main() {

	const c1 float64 = 3.14
	//const c1 float64 = math.Sin(3.14)  //常量不可再赋值,scala中val关键字
	fmt.Printf("%v,%T", c1, c1)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)

}
