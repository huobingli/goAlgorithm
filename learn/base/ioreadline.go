package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}
