package main

import (
	"fmt"
)

func main() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	fmt.Println()

	// 输出切片长度
	fmt.Println(len(slice))

	// 输出切片容量
	fmt.Println(cap(slice))
}
