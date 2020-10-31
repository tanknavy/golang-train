package main
import ("fmt")

//switch匹配变量值，bool表达式, v.(type)变量类型
func main(){
	populations := map[string]int{ //定义map[string]int
		"ca" : 29,
		"tx" : 22,
		"wa" : 18, //逗号必须
	}

	num := 50

	//if中operator还有 &&(and), ||(or)，区别bitwise
	if pop, ok := populations["co"]; ok {//条件先计算ok，再检测
		fmt.Println("pop is: ", pop)
	}
	//fmt.Println(pop)//上述变量pop仅仅在if中有效

	if num <= 100 && returnTrue(){
		fmt.Println("test ok")
	} else if num >100{
		fmt.Println("num > 100")
	} else {
		fmt.Println("nothing")
	}

	switch i:=2+3; i { //匹配一个，都没有就default，不需要break
	case 1,2,3: //可以多个
		fmt.Println("one to three")
	case 4,5,6:
		fmt.Println("four to six")
	default : //都不匹配
		fmt.Println("not one to six")
	}

	i:=9
	switch{ //匹配一个，都没有就default，
	case i <= 10: //可以多个
		fmt.Println("less than 10")
		fallthrough //即使满足这个条件也要往下看
	case i <= 20:
		fmt.Println("less than 20")
	default : //都不匹配
		fmt.Println("greater than 20")
	}

	var j interface{} = "8" //interface{}定义具有一组方法的类型
	switch j.(type) { //变量j的属性type switch
	case int:
		fmt.Println("int type")
		break //提前打断
	case string:
		fmt.Println("string type")
		break
	default:
		fmt.Println("other type")
		break
	}

}

func returnTrue()  bool{ //bool类型返回值
	fmt.Println("returning true function")
	return true
}