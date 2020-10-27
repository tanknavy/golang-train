//package main//测试中main包也可以
package myTest

import (
	_ "fmt" //前面加_暂时注销
)

//不用写主函数
//被测试函数
func addUpper(n int) int {
	res := 0
	for i:=1; i<=n; i++{
		res +=i
	}
	return res
}

//被测试函数
func sub(m int, n int) int {
	return m - n
}


//go run main.go
//go build -o test.exe
