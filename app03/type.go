package main
import ("fmt")

func main(){
	n:= 1==1 //定义变量n等于bool类型
	m:= 1==2
	fmt.Printf("%v,%T\n", n,n)
	fmt.Printf("%v,%T\n", m,m)
	
	a:=10
	b:=3
	//var c int64 = 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //3，类型int
	fmt.Println(a % b) //求余数
	fmt.Printf("%v,%T\n", a,a)
	//fmt.Println(a+c) //类型不符合

	s := "l love you" //字符串
	c := []byte(s) //字符串转字节数组
	d := byte('A') //字符转字节,A在ascii中是65
	r := 'a' //单引号表示chracter，类似java中char
	
	fmt.Printf("%v,%T\n",c,c)
	fmt.Printf("%v,%T\n",d,d)
	fmt.Printf("%v,%T\n",r,r)
	

}