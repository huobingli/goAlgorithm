package main

import (
	"fmt"
)

func main() {
	var sum *int
	sum = new(int) //分配空间
	fmt.Println(*sum) // 输出0
	*sum = 98
	fmt.Println(*sum) // 输出98

	type Student struct {
		name string
		age int
	 }
	 var s *Student
	 s = new(Student) //分配空间
	 s.name ="dequan"
	 fmt.Println(s)
}
