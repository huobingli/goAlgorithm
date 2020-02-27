package main
//import ."fmt"

// 交换函数
func swap(a, b *int) {
    // 取a指针的值, 赋给临时变量t
    t := *a
    // 取b指针的值, 赋给a指针指向的变量
    *a = *b
    // 将a指针的值赋给b指针指向的变量
    *b = t
}

func main() {
	// 位运算交换两值
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	println(a, b)

	swap(&a, &b)
	println(a, b)

	// 多重赋值
	var c int = 300
	var d int = 400
	d, c = c, d
	println(c, d)
}