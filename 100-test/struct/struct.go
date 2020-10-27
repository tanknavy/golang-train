package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------------------")
	p := Person{
		"bob",
		28,
	}
	s := Student{
		name: "tom",
		age: 19,
		getInfo: func(){
			//fmt.Println(name, age)
		},
	}

	p.getInfo() //都可以
	(&p).getInfo() //加不加&都可以
	s.getInfo()

}

type Person struct{
	name string
	age int
}
func (p *Person) getInfo(){
fmt.Println(p.name, p.age)
}


type Student struct{
	name string
	age int
	getInfo func()
}

type Monster struct{
	Name string
	Age int
	Skill []string
}


//go run main.go
//go build -o test.exe
