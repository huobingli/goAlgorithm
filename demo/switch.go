package main

import "fmt"

func main() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ': // 空格符会直接 break，返回 false // 和其他语言不一样
			fmt.Println("case 1")
		    //fallthrough // 返回 true
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) // true
	fmt.Println(isSpace(' ')) // false
}