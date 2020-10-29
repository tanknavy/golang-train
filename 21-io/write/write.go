package main
import (
	"fmt"
	"os" //File结构体，三种叫法，文件对象/指针/句柄
	"io" //使用下面的EOF错误, 
	"bufio" //带缓冲的io(文件较大时)
	"io/ioutil" //ioutil可以一次读取整个文件到内存(文件较小时),自带open/close
)

//os.OpenFile(name, flag, perm) //更一般的文件操作
//bufio.NewWriter() //带缓冲的写
//io.ioutil.WriteFile() //不带缓冲的写
func main(){
	fmt.Println("----------文件写操作----------------")
	fileName := "e:/tmp/t01.txt"
	
	//1. 打开文件，写入内容
	//1.1 不存在就创建
	//file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE, 0666)//只写，不存在就创建新文件
	//1.2 文件存在，先trunc再写入
	//file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_TRUNC, 0666)//先清除
	//1.3 文件存在，追加内容
	//file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0666)//追加
	//1.4 读写方法，追加
	file, err := os.OpenFile(fileName, os.O_RDWR | os.O_APPEND, 0666)//读写方式打开，追加内容
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return //打开出错，写不了
	}

	//最后.defer关闭资源
	defer file.Close() //函数退出时及时关闭file，否则内存泄露
	//读写模式打开时，可以先读取原来的文件内容
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //每次读取直到换行(一次读一行)，带缓存的，内容先写入到缓存的还没有落盘
		if err == io.EOF{ //读取到了文件末尾
			break
		}
		fmt.Print(str)
	}


	//2.写文件
	//2.1 bufio带缓冲区的写方式
	str := "高级软件工程师\r\n" //\n表示换行,有些编辑器不能正常处理\n，\r\n一起用
	writer := bufio.NewWriter(file) //默认缓冲大小4096bytes
	for i:=0;i<3;i++{
		writer.WriteString(str) //带缓存的，内容先写入到缓存的还没有落盘
	}
	writer.Flush() //将缓冲的数据写入文件中！
	fmt.Println("文件写入完成...")

	//2.2 ioutil读取整个文件，无需open/close,因为被封装到了
	content, err := ioutil.ReadFile(fileName)//返回[]byte字节切片
	
	if err != nil {
		fmt.Printf("read file err=%v",err)
	}
	//fmt.Println(content)
	fmt.Println(string(content))

	//3.关闭文件, 推荐使用defer，在函数结束之前调用关闭资源
	// err = file.Close()
	// if err != nil{
	// 	fmt.Println("close file err=",err)
	// }


}