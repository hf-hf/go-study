package main

import (
	"fmt"
)

func mutate(s string) string {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	return s
}

/**
我们尝试将字符串的第一个字符更改为'a'. 。
这是不允许的，因为字符串是不可变的，因此程序抛出错误 main.go:8: cannot assign to s[0]。
要解决字符串不能改变的问题，请将字符串转换为 rune 切片。然后，就可以愉快的更改了，然后返回新的字符串。
*/
func main() {
	h := "hello"
	fmt.Println(mutate(h))
}
