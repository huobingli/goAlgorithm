package main

import (
	"fmt"
)

func deferAdd(a, b int) func() int {

	fmt.Println(a, b)
	// 并没有执行闭包
	return func() int {
		return a + b
	}
}

func main() {
	// 此时返回的是匿名函数
	addFunc := deferAdd(1, 2)
	fmt.Println("-----------")
	// 这里才会真正执行加法操作
	fmt.Println(addFunc())
}
