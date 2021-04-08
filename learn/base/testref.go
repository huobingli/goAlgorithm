package main
import (
  "fmt"
)
func main() {
  m1 := make([]string, 1)
  m1[0] = "test"
  fmt.Println("调用 func1 前 m1 值:", m1)
  func1(m1)
  fmt.Println("调用 func1 后 m1 值:", m1)
}
func func1 (a []string) {
  a[0] = "val1"
  fmt.Println("func1中:", a)
}