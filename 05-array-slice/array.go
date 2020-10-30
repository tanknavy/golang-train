package main

import (
	"fmt"
)

//array & slice
//注意，在大部分语言中数组是引用类型，但golang中数据是值类型，slice是引用类型,go中指针,slice,map,channel是引用类型
func main() {

	// array：固定长度，声明[num]int, [5]int{},[...]int{1,2,3},要么前面[...]后面指定元素，要么[length]type，
	// slice：可变长度，声明[]int, []int{1,2,3},切片可以从现有数组创建,或者make([]int, length, capacity), 可以make，元素是指针，底层是array
	// array定义后就产生了空间，这是打印输出默认值，而slice没有make就没有空间，打印出来是[], 后面可以append数据
	grades := [...]int{1, 2, 3, 4, 5} //整型数组,数值默认0
	aa := make([]int, 3, 8)           // 创建slice，长度3，容量10
	//数组的地址就是第一个元素的地址，&grades等于&grades[0]，后一个元素地址是前一个地址+数组类型字节数
	fmt.Printf("值:%v,类型:%T,地址:%p\n", grades, grades, &grades) //[5]int
	fmt.Printf("make a array of capacity: %v, %v, %v\n", grades, len(grades), cap(grades))
	gg1 := grades
	//grades[0] = 11
	fmt.Printf("数组%v %v\n", grades, gg1) //数组，不一样，因为是值引用，其它语言比如java是引用

	grades2 := []int{1, 2, 3, 4, 5}         //创建slice,[]里面加个数或者..就是数组了
	fmt.Printf("%v %T\n", grades2, grades2) //[]int
	gg2 := grades2
	grades2[1] = 22
	fmt.Printf("slice%v %v\n", grades2, gg2) //切片，所以一样，因为是引用类型

	var students [3]string
	students[0] = "alex"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students: %v\n", len(students))
	fmt.Printf("students address: %v\n", &students)

	aa = append(aa, 1)                    //slice才能append追加元素
	aa = append(aa, []int{2, 3, 4, 5}...) //将slice逐个元素追加到slice
	fmt.Printf("aa slice of capacity: %v, %v, %v\n", aa, len(aa), cap(aa))

	bb := [5]int{} //数组是固定长度,
	bb[3] = 12     //数组可以index
	fmt.Printf("bb Array of capacity: %v, %v, %v\n", bb, len(bb), cap(bb))

	cc := make([]int, 3) //make创建slice,或者[]int{}指定，不加...
	//cc := []int{11,22,34} //slice是可变长度像list
	//cc[3] = 12 //定位时必须有，或者append
	cc = append(cc, []int{4, 5, 6, 7, 8}...) //追加在后面，前面默认为0
	fmt.Printf("cc slice of capacity: %v, %v, %v\n", cc, len(cc), cap(cc))

	//var p *[3]string = &students
	p := &students // 指针,:=类型推断更简洁
	fmt.Printf("%v, %T\n", p, *p)

	fmt.Println(grades[:2])  //数据切片，前n个，不包含n
	fmt.Println(grades[3:])  //数据切片，从第4个开始
	fmt.Println(grades[2:4]) //数据切片，从第4个开始

}
