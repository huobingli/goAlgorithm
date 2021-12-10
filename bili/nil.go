package main

import "fmt"

type MyError struct{}

func (hello *MyError) Error2() string {
	return "1"
}

func test() error2 {
	var p *MyError = nil
	/*
	   //Check err and Set p
	   if bad() {
	       p = ErrBad
	   }
	*/
	return p
}

func main() {
	err := test()
	fmt.Print(err)
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is NOT nil")
	}
}
