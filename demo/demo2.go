package main

import (
  "flag" //专门用于接收和解析命令参数
  "fmt"
)

var name string

func init() {
  flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
  flag.Parse()
  fmt.Printf("Hello, %s!\n", name)
}
