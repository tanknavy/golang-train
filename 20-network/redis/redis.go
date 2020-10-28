package main

//redis-server.exe redis.windows.conf
//redis-cli.exe -h 127.0.0.1 -p 6379
import (
	"fmt"

	"github.com/garyburd/redigo/redis" //外部包全路径，使用go get安装在gopath里面
)

//golang operate Redis，Do必要时会转换为相应类型
func main() {
	fmt.Println("------------------------------")

	//1.连接Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis failed--->", err)
	}
	fmt.Println("conn succ...", conn)

	//2.关闭连接
	defer conn.Close()

	//3.Redis Set/HSet
	_, err = conn.Do("Set", "go1", 985)
	if err != nil {
		fmt.Println("set key failed\n", err)
		return
	}

	//4.Redis get，使用自带的类型转换redis.Int, redis.Strings
	r, err := redis.Int(conn.Do("Get", "go1")) //返回的r是interface{}类型，需要对应的转换,不要自己强转，
	if err != nil {
		fmt.Println("get key failed\n", err)
		return
	}

	fmt.Println(r)

	//5.HSet,设置类似字典对象，逐个元素赋值, MSet,MGet批量操作多个元素
	_, err = conn.Do("HSet", "user01", "name", "tom")
	if err != nil {
		fmt.Println("HSet key set failed\n", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("HSet key set failed\n", err)
		return
	}

	hs01, err01 := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HSet key get failed\n", err01)
		return
	}
	fmt.Println(hs01)

	hs02, err02 := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HSet key get failed\n", err02)
		return
	}
	fmt.Println(hs02)

	//6. HMSet, HMGet一次写读一个map
	_, err = conn.Do("HMSet", "user02", "name", "Jack", "age", 28)//一次set多个字段
	if err != nil {
		fmt.Println("HSet key set failed\n", err)
		return
	}

	//redis.Strings
	hs03, err03 := redis.Strings(conn.Do("HMGet", "user02","name", "age"))//一次get多个字段,返回一个切片
	if err != nil {
		fmt.Println("HSet key get failed\n", err03)
		return
	}

	fmt.Println(hs03)
	for i,v := range hs03 {
		//fmt.Printf("hs03:%s\n",v)
		fmt.Printf("hs03[%d]=%s\n",i,v)
	}
	

}

//go run main.go
//go build -o test.exe
