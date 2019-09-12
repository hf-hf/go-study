package main

import (
	"fmt"
	"unicode/utf8"
)

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

/**
utf8 package的func RuneCountInString(s string) (n int) 函数用于查找字符串的长度。
此方法将字符串作为参数，并返回他的的 runes 数。
*/
func main() {
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)
}
