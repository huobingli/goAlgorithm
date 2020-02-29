package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	// *skillParam 命令行参数传入
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
