package main

import (
	"fmt"
)

func main() {
	var n = []int{1, 2, 3, 4}
	for key, value := range n[1:2] {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	var str = "hello world"
	for key, value := range str[1:2] {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
