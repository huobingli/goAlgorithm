// 传值 传指针

package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a + 1 // 我们改变了a的值
	return a  //返回一个新值
}

// 传入地址
func add2(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}

func main() {
	x := 3
	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add1(x) //调用add1(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"

	x2 := add2(&x) // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x2) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
}
