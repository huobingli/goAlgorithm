package main

import "fmt"

var a [3]int // 定义三个整数的数组
func main() {

	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d", i, v)
	}

	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d", v)
	}

	fmt.Println()
	fmt.Println("-------------------------")
	//var q [3]int = [3]int{1, 2, 3}
	//fmt.Println(q[2])

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	q := [...]int{1, 2, 3}
	fmt.Printf("%T", q) // "[3]int"
}
