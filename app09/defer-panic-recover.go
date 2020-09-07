package main
import ("fmt";"io/ioutil";"log";"net/http")
//defer, panic, recover

func main(){
	defer fmt.Println("start") 
	defer fmt.Println("middle")
	defer fmt.Println("end")//defer使用LIFO方式决定执行顺序

	a := "aa"
	defer fmt.Println("----defered a is :",a) //注意，这时还是aa
	a = "bb" //改变a的值
	fmt.Println("----current a is :", a) //输出bb

	httpGet() //执行函数
	
	fmt.Println("----------------------------------")
	c,d := 2, 1
	res := c / d
	fmt.Println(res)
	
	fmt.Println("-------------defer/panic-recover---------------------")
	//panic在最后，但是使用了defer的recover可以捕捉到panic, 这时全部defer又会在panic后面依次发生
	//func(){ //匿名函数，和下面defer修饰比较
	defer func(){ //匿名，延时函数
		//使用defer+recover类似hook,捕捉到左右后可以继续全部defer动作
		if err := recover(); err != nil{ //recover在panic后面,有点类似try/catch的finally
			fmt.Println("with defer/recover, defer can still continue to do")
			log.Println("Error:", err)
	}
	}() //()调用这个匿名函数

	panic("something bad happend!")//最后抛出的错误，程序退出，panic比defer靠后(除非defer+recover)
	fmt.Println("the end") //在panic后面，不可到达代码

}


func httpGet(){
	res,err := http.Get("http://www.google.com/rebots.txt")
	if err != nil { //错误不为空
		log.Fatal(err)
	}
	defer res.Body.Close() //延时关闭资源
	rebots, err := ioutil.ReadAll(res.Body)//读取http返回体
	//res.Body.Close() //关闭资源

	if err != nil {//错误不为空
		log.Fatal(err)
	}
	fmt.Printf("%s", rebots) //答应http响应

}

