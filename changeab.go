package main
//import ."fmt"

func main() {
	// 位运算交换两值
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	println(a, b)

	// 多重赋值
	var c int = 300
	var d int = 400
	d, c = c, d
	println(c, d)
}