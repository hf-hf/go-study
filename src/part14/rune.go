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
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

/**
rune是go的内置类型，它是int32的别名。rune表示Go中的Unicode代码点。
代码点占用多少字节并不重要，它可以用 rune 表示。让我们修改上面的程序来使用rune来输出字符。
*/
func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
