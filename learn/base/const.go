package main

import "fmt"

//第一个外其它的常量右边的初始化表达式都可以省略，
//如果省略初始化表达式则表示使用前面常量的初始化表达式，
//对应的常量类型也是一样的
const (
	a = 1
	b
	c = 2
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
