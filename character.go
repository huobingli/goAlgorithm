package main

import ."fmt"

func main() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}