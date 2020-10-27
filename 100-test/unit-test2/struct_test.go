package test2
import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T){
	//1.创建一个实例
	monster := &Monster{ //值传递还是引用，这里是引用所以可以不用&
	//monster := Monster{ //可以不使用地址指针，因为sturct是引用类型
		Name : "red cow",
		Age : 38,
		Skill : "power",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("Store()错误，expectd=%v, result=%v", true, res)
	}
	t.Logf("Store()测试ok!")
}

func TestRestore(t *testing.T){
	//反序列化
	//还是要先创建一个实例，但是不需要指定字段的值
	var monster = &Monster{} //变量, 空值因为要求是指针类型，但是这里也可以不用写&,因为struct就是引用类型
	//var monster Monster //变量上面，和这里都可以
	fmt.Println(monster.Name, monster.Age) //测试看看

	res := monster.Restore()
	if !res {
		t.Fatalf("Restore()错误，expectd=%v, result=%v", true, res)
	}
	t.Logf("Restore()测试ok!")

	if monster.Name != "red cow" { //判断字段值
		t.Fatalf("Restore()错误，expectd=%v, result=%v", true, res)
	}
	fmt.Printf("返回Name:=%s",monster.Name)
	t.Logf("Restore()测试ok!")
}