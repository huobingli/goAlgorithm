package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)

	fmt.Printf("road:, %d\n", *r)
}
func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}
func main() {
	fmt.Printf("---begin---\n")
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)

		fmt.Println(time.Second)
		runtime.GC()
	}

	fmt.Printf("---end---\n")
}
