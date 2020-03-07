package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello go")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	// ReadByte
	c, err := r.ReadByte()
	fmt.Println(string(c), err)

	// ReadBytes
	var delim byte = ' '
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}
