package main

import (
	"fmt"
)

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

/**
以Unicode字符方式遍历时，每个字符的类型是rune，而不是byte。rune类型在go语言中占用四个字节。
在go语言中支持两个字符类型，一个是byte(实际上是uint8的别名)，
代表UTF-8字符串的单个字节的值，另一个是rune，代表单个Unicode字符。
*/
func main() {
	name := "Señor"
	for index, test := range name {
		fmt.Printf("%c starts at byte %d\n", test, index)
	}
	printCharsAndBytes(name)
}
