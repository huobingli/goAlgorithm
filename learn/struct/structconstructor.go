package main

import (
	"fmt"
)

type Cat struct {
	Color string
	Name  string
}
type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func main() {
	// 1
	//var greyCat = new(BlackCat)

	// 2
	//var greyCat = NewBlackCat("grey")
	//var greyCat *BlackCat = NewBlackCat("grey")
	//greyCat := NewBlackCat("grey")

	// 3
	var greyCat BlackCat
	greyCat.Name = "nini"
	greyCat.Color = "red"

	fmt.Println(greyCat.Name)
	fmt.Println(greyCat.Color)
}
