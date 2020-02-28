package main

import (
	"fmt"
)

/**
冒泡排序
*/
func main() {
	arr := [...]int{21, 32, 12, 33, 34, 34, 87, 24}
	var n = len(arr)
	fmt.Println("--------没排序前--------\n", arr)
	for i := 0; i <= n-1; i++ {
		fmt.Println("--------第", i+1, "次冒泡--------")
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
			fmt.Println(arr)
		}
	}
	fmt.Println("--------最终结果--------\n", arr)
}
