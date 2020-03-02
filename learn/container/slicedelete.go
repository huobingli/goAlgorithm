package main

import (
	"fmt"
)

func main() {
	//a := [3]int
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a = append(a[:0], a[1:]...) // 删除开头1个元素

	N := 3

	//for _, v := range a {
	for i := 0; i < len(a); i++{
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()

	a = append(a[:0], a[1:]...) // 删除开头1个元素

	for i := 0; i < len(a); i++{
		fmt.Printf("%d ", a[i])
	}

	fmt.Println()

	a = append(a[:0], a[N:]...) // 删除开头N个元素

	for i := 0; i < len(a); i++{
		fmt.Printf("%d ", a[i])
	}

	fmt.Println()
}
