package main
import (
	"fmt"
	"os" //File结构体，三种叫法，文件对象/指针/句柄
	"io" //使用下面的EOF错误, 
	"bufio" //带缓冲的io(文件较大时)
	"io/ioutil" //ioutil可以一次读取整个文件到内存(文件较小时),自带open/close
)

//流：输入流，输出流
//文件的打开和关闭
//文件的读写,判断
func main(){
	fmt.Println("----------文件读操作----------------")

	//1.打开文件
	file, err := os.Open("e:/tmp/test.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	fmt.Printf("file=%v\n",file) //0xc000076780

	//defer关闭资源
	defer file.Close() //函数退出时及时关闭file，否则内存泄露

	//2.读取文件
	//2.1 bufio带缓冲区的方式
	reader := bufio.NewReader(file) //默认缓冲大小4096bytes
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //参数是delimiter, 读到一个delimiter(\n换行)就结束一次，就是一行一行的读取
		if err == io.EOF { //io.EOF表示文件的莫问
			break
		}
		fmt.Print(str)//读取的时候已经有换行了

	}
	fmt.Println("文件读取完成...")

	//2.2 ioutil读取整个文件，无需open/close,因为被封装到了
	content, err := ioutil.ReadFile("e:/tmp/test.txt")//返回[]byte字节切片
	if err != nil {
		fmt.Printf("read file err=%v",err)
	}
	fmt.Println(content)
	fmt.Println(string(content))

	//3.关闭文件, 推荐使用defer，在函数结束之前调用关闭资源
	// err = file.Close()
	// if err != nil{
	// 	fmt.Println("close file err=",err)
	// }
	
	_, err = os.Stat("e:/tmp/test.txt")
	if err == nil{
		fmt.Println("文件或目录存在")
	}
	if os.IsNotExist(err){ //判断错误是否是文件/目录不存在
		fmt.Println("文件不存在")
	}

}