package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //外部包全路径，使用go get安装在gopath里面
)

//redis connection pool
//全局变量，指针类型
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init(){ //初始化函数
	
	pool = &redis.Pool{ //新建一个连接池
		MaxIdle: 8,
		MaxActive: 0,
		IdleTimeout: 100,
		Dial :func()(redis.Conn, error){
			return redis.Dial("tcp", "127.0.0.1:6379") //连接池如何连接redis server
		},
	}

}


func main() {
	fmt.Println("------------------------------")
	// 从pool中取出一个连接
	conn := pool.Get()
	
	defer conn.Close() //延时关闭连接

	res,err := redis.Int(conn.Do("Get","go1"))
	if err != nil {
		fmt.Println("get error", err)
	}
	fmt.Println(res)

}

//go run main.go
//go build -o test.exe
