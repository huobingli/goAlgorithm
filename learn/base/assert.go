package main

import (
	"fmt"
)

func main() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ",", ok)
}
