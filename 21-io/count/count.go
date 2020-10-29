package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

type CharCount struct{ //也可以用map,但struct变量定义后有每个属性默认值，而map需要初始化每个key，输出时还无需
	ChCount int //字母个数
	NumCount int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其它字符个数
}

var countMap map[string]int //使用map, 输出也无需
//var countMap = make(map[string]int,4)


//统计一个文件中,英文字符，汉字，数字，空格的个数
func main() {
	fmt.Println("------------------------------")
	fileName := "e:/tmp/t01.txt"

	//1.打开文件，创建一个reader
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}

	//2.关闭资源
	defer file.Close()

	//3.结果实例，用struct或者map
	var count CharCount //有默认值
	countMap = map[string]int{"chCount":0, "NumCount":0, "SpaceCount":0, "OtherCount":0} //map初始化

	//4.对于每行遍历，统计将结果保存到一个结构体中
	reader := bufio.NewReader(file)

	//5.循环读取文件内容
	for {
		line, err := reader.ReadString('\n') //\n换行，一次读一行
		//line, err := reader.ReadBytes('\n') //\n换行，一次读一行
		if err == io.EOF{ //到了文件末尾
			break
		}
		//遍历读取到的一行, 是个字符串, 字符串本质就是字节数组[]byte, range时就是一个个byte, 字符比较大小就是ascii码比较
		for _, v := range line { //循环字符串
			//fmt.Println(v)
			//switch v { //里面是bool，这里是byte，不匹配
			switch { // 不用变量，直接将switch当做if else分支结构
				case v >= 'a' && v <= 'z': // ascii编码比较大小
					//count.ChCount ++	
					fallthrough //关键字，穿透，表示直接使用下一个的count.ChCount ++
				case v >= 'A' && v <= 'Z': // ascii编码比较大小
					count.ChCount ++
					countMap["chCount"] = countMap["chCount"] + 1
				case v == ' ' || v == '\t': // 空格或者ta
					count.SpaceCount ++
					countMap["SpaceCount"] = countMap["SpaceCount"] + 1
				case v >= '0' && v <= '9': //在
					count.NumCount ++
					countMap["NumCount"] = countMap["NumCount"] + 1
				default:
					count.OtherCount ++	//注意，\r\n看不见但是也算是字符
					countMap["OtherCount"] = countMap["OtherCount"] + 1
			}
		}
	}


	//6.最终结果
	fmt.Printf("字符个数=%v, 数字个数=%v, 空格个数=%v, 其它个数=%v\n",count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
	fmt.Println(countMap)
}

//go run main.go
//go build -o test.exe //编译成可执行文件
//go test -v //使用testing框架测试_test.go结尾的文件
