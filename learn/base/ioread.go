package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello world")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}
