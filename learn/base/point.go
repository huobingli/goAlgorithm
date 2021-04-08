package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ptr *int
	fmt.Println(ptr)

	a := 100
	ptr = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr1 := new(int)
	*ptr1 = 100

	fmt.Println(*ptr1)

	b := 100
	c := 101

	c, b = b, c

	fmt.Println(b, c)

	/*
		uintptr 是 Go 内置的可用于存储指针的整型，而整型是可以进行数学运算的！
		因此，将 unsafe.Pointer 转化为 uintptr 类型后，就可以让本不具备运算能力的指针具备了指针运算能力：
	*/
	arr := [3]int{1, 2, 3}
	ap := &arr

	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	*sp += 3
	fmt.Println(arr)
}
