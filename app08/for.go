package main
import ("fmt")

func main(){
	for i:= 0; i<5; i++ { //类似java的for,没有(),可类型推断
		fmt.Println(i)
	}

	i := 0
	for i<5 { //可简写
		fmt.Println(i)
		i++
	}

	for i,j := 0,0; i<5; i,j = i+1,j+2 { //可同时写两个变量
		fmt.Println(i,j)
	}
	//fmt.Println(i) //i只在loop的作用域
	
	
	fmt.Println("---------------------")
	loop: //循环labels
	for i:=1;i<=3;i++{
		for j:=1;j<=3;j++{
			fmt.Println(i * j)
			if i * j >= 6 {
				//break //中断内部的for循环
				break loop //中断整个loop范围的循环
			}
			
		}
	}

	populations := map[string]int{ //定义map[string]int
		"ca" : 29,
		"tx" : 22,
		"wa" : 18, //逗号必须
	}

	fmt.Println("---------------------")
	aa := []int{11,22,33} 
	for k,v := range aa{ //range类似python的enumerate
		fmt.Println(k,v)
	}

	for k,v := range populations{ //range迭代map
		fmt.Println(k,v)
	}

	//go要求每个变量都要使用，不使用时可用_指代
	for _,v:= range populations{ //range迭代map
		fmt.Println(v)
	}

	fmt.Println("---------------------")
	s := "hello go!"
	for k,v := range s{ //迭代字符串
		fmt.Println(k,v)
	}


}
