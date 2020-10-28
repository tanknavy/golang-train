package main

import (
	"sort"
	"fmt"
)

//map如果无序，如何按照map的key顺序进行排序输出？
//1.现将map的key放入到切片中
//2.对切片排序
//3.遍历切片，然后按照key来输出map的值
func main() {
	m1 := make(map[int]string, 10)
	m1[9] = "giggs"
	m1[2] = "bob"
	m1[3] = "carl"
	m1[1] = "adm"

	fmt.Println(m1)
	
	var keys []int
	var arr [4]int
	for k := range m1 { //range是关键字，不是函数
		keys = append(keys, k)
		//arr = append(arr, k)
	}

	//排序的包sort
	sort.Ints(keys) //升序排序
	fmt.Println(keys)

	for _, k := range keys{ //下标，值
		fmt.Printf("map[%d]=%s\n",k, m1[k])
	}


}
