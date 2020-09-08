package main

import (
	"fmt"
	"time"
	//"sync"
)

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct{
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50) //logEntry类型的chan, 缓冲大小50
var doneCh = make(chan struct{}) //zero channel
//var wg = sync.WaitGroup{}

func main() {
	//wg.Add(1)
	go logger() //goroutine协程
	defer func(){
		close(logCh) //当主线程退出前，关闭goroutine
	}()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"} //输入消息
	logCh <- logEntry{time.Now(), logWarning, "App is shuting down"} //输入消息
	//close(logCh) //避免deadlock,
	time.Sleep(100 * time.Millisecond)//不想用Sleep就用waitGroup让主线程等待，
	doneCh <- struct{}{} //没有任何field的struct，常见方式
	//wg.Wait()
}

func logger(){
	//for entry := range logCh{ //for循环读出，range迭代logCh信道
	for {
		// choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operation is ready. 
		// If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statement will be a channel operation.
		// select监控多个channel直到一个ready，如果多个ready随机选择一个
		select{ //循环监控多个channel
		case entry := <-logCh: //是log信道时
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2012-09-06T15:04:05"), entry.severity, entry.message)
		case <- doneCh: //当doneCh有值时
			//fmt.Println(stop)
			break
		}
		
	}
	//wg.Done()
}