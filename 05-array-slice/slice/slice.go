package main
import (
	"fmt"
)

func main(){
	var arr [5]int = [...]int{1,2,3,4,5}
	var s1 []int = arr[1:3] //slice从现有array创建
	var s2 []int //只定义是[]，可以make开辟空间，或者append
	var s3 []int = make([]int,2,4) //开辟了空间，

	fmt.Println(len(s1),cap(s1))
	s2 = append(s2,1)
	s3 = append(s3,11,12)//从后加
	s3 = append(s3,22,23,24,25,26)//从后加,自动扩展
	fmt.Println(s3[0:])
	fmt.Println(len(s3),cap(s3))

	fmt.Println("arr=",arr)
	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
	fmt.Println("s3=",s3)

}