package main

//Go中首字母大写表示global
import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	//1.interface和struct
	//interface引用struct,可以调用实现的方法，
	//struct保存状态，interface持有各种方法
	var w Writer = ConsoleWriter{} //interface引用struct
	w.Write([]byte("Hello Go!"))
	fmt.Println(w.(ConsoleWriter))

	//2----使用interface/ int type实现计数器-----------
	//为什么需要struct,因为需要保存转态，函数不能保存状态
	ic := IntCounter(0)       //初始化一个int类型type
	var inc Incrementer = &ic //int类型type的地址作为
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Incre())
	}
	//interface方法中使用struct的地址作为输入参数，所以这样可以读取struct地址
	fmt.Println(inc.(*IntCounter)) //interface中的struct的地址？

	//3.interface的嵌套(继承)，多个方法实现
	var wc WriterCloser = NewBufferedWriterCloser() //返回*BufferedWriterCloser一个struc的指针
	//wc.Write([]byte{"Hello, this is a test"}) //错误，元素要字符
	//wc.Write([]byte{'h','e','y'})//正确
	wc.Write([]byte("Hello, this is a test")) //字符串转字节slice
	wc.Close()

	fmt.Println(wc.(Writer))          //wc中有Writer接口，Closer接口，还有BufferedWriterCloser
	bwc := wc.(*BufferedWriterCloser) //wc接口方法聚合了struct的地址作为指针的输入类型
	fmt.Println("interface中对sturct的引用：", bwc)

	//5.上述接口对象查看聚合了io.Reader接口
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("转换失败")
	}

	//6.查看接口类型
	var i interface{} = 0       //使用接口作为引用
	fmt.Printf("%v %T\n", i, i) //0,int
	switch i.(type) {
	case int:
		fmt.Println("is integer")
	case string:
		fmt.Println("is string")
	default:
		fmt.Println("don't now what it is ")
	}
}

//1.接口不描述数据而是描述行为！
type Writer interface { //定义接口，方法的集合，没有具体实现
	Write([]byte) (int, error) //方法(输入) (输出)
}

type ConsoleWriter struct{} //定义结构体(类似class)

//struct绑定的method实现interface的方法
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Closer interface {
	Close() error
}

//3---嵌套两个接口，类似继承
type WriterCloser interface {
	Writer //embedding嵌套接口,类似继承
	Closer //embedding嵌套接口,类似继承
}

//struct类似实体对象，暂时只有属性
type BufferedWriterCloser struct {
	buffer *bytes.Buffer //有read/write方法的可变长度字节的buffer
}

//引用BufferedWriterCloser结构体，实现interface的Write方法
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) //将data数据写到buffer缓冲区
	fmt.Println("写数据", n, "个字符到缓冲区")
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)       //8字节slice缓冲区
	for bwc.buffer.Len() > 8 { //for条件每8字节输出一次
		_, err := bwc.buffer.Read(v) //从缓冲区抽取v个字节的数据
		if err != nil {
			return 0, err
		}
		fmt.Println("Write阶段每8字节长度输出")
		_, err = fmt.Println(string(v)) //输出
		if err != nil {
			return 0, err
		}
	}
	return n, nil //最终返回
}

//引用BufferedWriterCloser结构体，实现interface的Close方法
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 { //关闭时如果还有字符
		data := bwc.buffer.Next(8)
		fmt.Println("Closer阶段输出最后一小段小于8字节的")
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser { //返回指针
	return &BufferedWriterCloser{ //初始化struct, 返回地址
		buffer: bytes.NewBuffer([]byte{}), //默认输入空字节slice
	}
}

//2.计数器的interface,type, func------------------------------
type Incrementer interface {
	Incre() int
}

type IntCounter int //使用时初始化IntCounter(0)

//为什么要使用指针？如果不使用每次传入的ic的值都会复制一份，每次ic都等于1
func (ic *IntCounter) Incre() int { //这里不像是func绑定了type，而是type的地址
	*ic++
	return int(*ic)
}

//----------------------------------------
func test01() int { //方法不能保存状态!!!
	num := 0
	num++
	fmt.Println("current num is:", num)
	return num
}
