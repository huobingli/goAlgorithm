package main

import "fmt"

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
	rawPrint(slist)
}

// arg是不定数量的参数, 变量arg是一个int的slice：
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	print(1, 2, 3)
	myfunc(1, 2, 3)
}
