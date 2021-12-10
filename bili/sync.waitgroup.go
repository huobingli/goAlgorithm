package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10000; i++ {
		x = x + 1
	}
	fmt.Println(x)
	wg.Done()
}

func main() {
	wg.Add(1)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
