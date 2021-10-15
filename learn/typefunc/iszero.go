package main

import (
	"fmt"
)

// 将int定义为MyInt类型
type MyInt int

// 为MyInt添加IsZero()方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.Add(2))

	fmt.Println(int(^uint(0) >> 1))

	// x := 1234
	// s := strconv.Itoa(x)

	aa := "123321"
	bb := "123321"
	fmt.Print(aa[0] == bb[0])
	// if s == s[::-1]{
	//     print(1)
	// } else {
	//     print(0)
	// }
}
