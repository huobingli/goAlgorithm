package main

import (
	"fmt"
)

func main() {

	//JLoop:
	for j := 0; j < 5; j++ {
		JLoop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Printf("%d ", i)
		}

		fmt.Println(111)
	}
	
	fmt.Println(222)
}
