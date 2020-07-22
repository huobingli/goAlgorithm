// 多路复用的能力，使得 goroutine 可以同时等待多个 channel 的读写操作

package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	// 循环向 channel 发送消息
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}
func receive(ch <-chan int) {
	val := <-ch
	fmt.Println("receive:", val)
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 0)
	go receive(ch2)
	// 主 goroutine 休眠 1s，保证调度成功
	time.Sleep(time.Second)
	for {
		select {
		case val := <-ch1: // 从 ch1 读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case ch2 <- 2: // 使用 ch2 发送消息
			fmt.Println("send value by ch2")
		case <-time.After(2 * time.Second): // 超时设置
			fmt.Println("Time out")
			return
		}
	}
}
