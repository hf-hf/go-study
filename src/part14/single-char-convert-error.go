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
程序的第31行，我们输出 Señor 字符，它输出的却是错误的 S e Ã ± o r。
为什么上面的 Hello World 没问题，而下面的 Señor 有问题呢？原因是 ñ 的Unicode编码为U+00F1，
其 UTF-8 encoding 占用2字节，分别是c3和b1。我们输出字符时，是假设每个代码点将是一个字节长，
这是错误的。在UTF-8编码中，一个代码点可以占用超过一个字节。我们该怎么解决这个问题呢？这就是 rune 派上了用场。
*/
func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
