package main

import (
	"fmt"
)

func main() {
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}

	for _, value := range slice {
		for index2, value2 := range value {
			fmt.Printf("Index: %d Value: %d\n", index2, value2)
		}
	}
}