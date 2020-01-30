
package main

import (."fmt"
		."flag"
)

var block = "package"

func main() {
  block := "function"
  {
    block := "inner"
    Printf("The block is %s.\n", block)
  }
  Printf("The block is %s.\n", block)
}