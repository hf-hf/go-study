package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

/**
Go中的字符串是Unicode兼容的，并且是UTF-8编码。
由于字符串是字节片，所以可以访问字符串的每个字节。
*/
func main() {
	name := "Hello World"
	printBytes(name)
}

/**
上面的程序第八行中，len(s) 返回字符串中的字节数，我们使用for循环输出这些字节的十六进制。
％x 是十六进制的格式说明符。上述程序输出 48 65 6c 6c 6f 20 57 6f 72 6c 64。这些是“Hello
World”的 Unicode UT8-encoded 编码值。
需要对Unicode和UTF-8有基本的了解才能更好地理解字符串。
我建议阅读https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/ 来了解什么是Unicode和UTF-8。
*/
