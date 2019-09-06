package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

/**
%c 格式说明符用于输出字符串的字符
*/
func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
