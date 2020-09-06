package main
import ("fmt")

func main(){
	
	const c1 float64 = 3.14
	//const c1 float64 = math.Sin(3.14)  //常量不可再赋值,scala中val关键字
	fmt.Printf("%v,%T",c1,c1)
	

}