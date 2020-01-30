package main

import (
  // 需在此处添加代码。[1]
  "flag" //专门用于接收和解析命令参数
  "fmt"
  "os"
)

var name string

func init() {
    //flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
    flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
  	flag.CommandLine.Usage = func() { 
	  	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question") 
	  	flag.PrintDefaults()
    }
    flag.StringVar(&name, "name", "everyone", "The greeting object.")  
}

func main() {
	// flag.Usage = func() { 
  //   fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question") 
  //   flag.PrintDefaults()
	// }
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
