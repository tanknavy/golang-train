package test2
import (
	"fmt"
	"encoding/json" //json序列化
	"io/ioutil" //io工具，读写文件
)

type Monster struct{
	Name string
	Age int
	Skill string
}

//绑定方法，序列化
func(m *Monster) Store() bool{
	//1.序列化
	data, err := json.Marshal(m) //返回[]byte
	if err != nil{
		fmt.Println("marshal err=", err)
		return false
	}
	//2.序列化后保存到文件
	filePath := "e:/tmp/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil{
		fmt.Println("write file err=", err)
		return false
	}
	return true
}

//绑定方法，反序列化
func(m *Monster) Restore() bool{
	//1.从文件中读取序列化的字符串
	filePath := "e:/tmp/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("read file err=", err)
		return false
	}

	//2.使用读取到的数据[]byte，反徐序列化
	err = json.Unmarshal(data, m)
	if err != nil{
		fmt.Println("unmarshal err=", err)
		return false
	}
	return true
}