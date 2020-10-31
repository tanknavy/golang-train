package main

import (
	"fmt"
	"sort"
	"math/rand"
)

//最经典的面向interface编程
//1.结构体的切片排序，func Sort(data Interface)，
//Interface：A type, typically a collection，一个type, 一般是个集合，包含Len,Less,Swap三个方法
type Hero struct {
	Name  string
	Age int
}

//2.声明结构体的slice
type heroSlice []Hero //定义类型，除了struct和interface,type可以定义各种自定义类型

//3.对这个集合实现Interface接口的三个方法
func (heros heroSlice) Len() int{
	return len(heros)
}

func (heros heroSlice) Less(i,j int) bool{ //怎么排序
	//return heros[i].Name > heros[j].Name //是否i应该排在j前面
	return heros[i].Age <= heros[j].Age //是否i应该排在j前面
}

func (heros heroSlice) Swap(i,j int) {
	// tmp := heros[i]
	// heros[i] = heros[j]
	// heros[j] = tmp
	heros[i],heros[j] = heros[j],heros[i] //更简洁
}

func main() {
	
	//1.一个数组/切片的排序
	var intArray = [...]int{0,-1,99,7,10}
	var intSlice = []int{0, -1, 99, 7, 10}
	
	//2.自己写排序算法，或者sort包
	sort.Ints(intArray[:]) //需要传入slice，这样将一个数组当做slice(每个元素是指针)的底层数组
	fmt.Println(intArray)
	//sort.Ints(intSlice) //原地排序，切片传入的时候不需要加地址，它本来就是引用类型的
	//附加Interface到[]int进行升序排,"Interface A type, typically a collection"
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)

	//3.如何对结构体切片进行排序？
	//func Sort(data Interface)，实现了Interface结构的变量，包含Len,Less,Swap三个方法,默认quickSort
	fmt.Println("-------Hero Rank---------------")
	var heros heroSlice
	
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d",rand.Intn(108)), //Sprintf组合一个格式化的字符串
			Age : rand.Intn(32) + 18,
		}
		heros = append(heros, hero) //放入
	}

	fmt.Println(heros) //排序前
	sort.Sort(heros) //结构体切片排序
	fmt.Println(heros) //排序后

}


