package main
import (
	"fmt"
)

//slice底层用array，从array中切片时slice中元素是指针类型指向array中的元素
//slice使用make产生时，底层array程序员看不见
func main(){
	var arr [5]int = [...]int{1,2,3,4,5}
	var s1 []int = arr[1:3] //slice从现有array创建
	var s2 []int //只定义是[]，可以make开辟空间，或者append
	var s3 []int = make([]int,2,4) //开辟了空间，
	var arr2 [4]int //4个元素长度数组，默认全为0
	var s4 = arr2[0:4] //变成一个slice，len=4,cap=4
	var s5 []int //默认capaticy是0
	//var s6 = s5[0:4] //s5 cap=0, 越界了！

	fmt.Println(len(s1),cap(s1))
	s2 = append(s2,1)
	s3 = append(s3,11,12)//从后加
	s3 = append(s3,22,23,24,25,26)//从后追加,自动扩展
	fmt.Println(s3[0:])
	fmt.Println(len(s3),cap(s3))

	fmt.Println("arr=",arr)
	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
	fmt.Println("s3=",s3)
	fmt.Println("s4=", s4, len(s4),cap(s4))//
	fmt.Println("s5=", s5, len(s5),cap(s5))//
	//fmt.Println("s6=", s6, len(s6),cap(s6))//
}