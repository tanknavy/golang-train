package main
import ("fmt")

func main() {
	
	aa := make([]int, 3,8) // 创建slice，长度3，容量10
	grades := [...]int{1,2,3,4,5} //整型数组,默认0
	fmt.Printf("%v\n", grades)

	var students [3]string
	students[0] = "alex"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students: %v\n", len(students))
	fmt.Printf("students address: %v\n", &students)
	
	aa = append(aa,1)//追加元素
	aa = append(aa, []int{2,3,4,5}...) //将数组逐个元素追加到数组
	fmt.Printf("make a array of capacity: %v, %v\n", aa,cap(aa))
	
	//var p *[3]string = &students
	p := &students // 指针,:=类型推断更简洁
	fmt.Printf("%v, %T\n", p, *p)

	fmt.Println(grades[:2])//数据切片，前n个，不包含n
	fmt.Println(grades[3:]) //数据切片，从第4个开始
	fmt.Println(grades[2:4]) //数据切片，从第4个开始
	
}
