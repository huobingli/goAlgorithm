package main

import (
	"fmt"
)

func main() {
	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	fmt.Println(aSlice)

	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	fmt.Println(aSlice)
	aSlice = array[:] // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素
	fmt.Println(aSlice)

	// 从slice中获取slice
	aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4，cap=7
	fmt.Println(aSlice)
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	fmt.Println(bSlice)
	bSlice = aSlice[:3] // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	fmt.Println(bSlice)
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	fmt.Println(bSlice)
	bSlice = aSlice[:] // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(bSlice)

	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值
	bSlice[1] = 'z'
	fmt.Println(aSlice)
	fmt.Println(array)

	println("-----------")

	var slice1 []byte
	slice1 = append(aSlice, 0)
	println(aSlice)
	println(slice1)
	fmt.Println(slice1)
}
