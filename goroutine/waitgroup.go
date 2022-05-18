package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
// 主程序会直接退出
func main() {
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
	wg.Wait()
}

func doIt(workerID int) {
	defer wg.Done()
	fmt.Printf("[%v] is running\n", workerID)
	time.Sleep(3 * time.Second)        // 模拟 goroutine 正在执行
	fmt.Printf("[%v] is done\n", workerID)
}