package main

import (
	"fmt"
)

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

/**
将字符串转换为切片后，即可对字符串进行更改，然后返回新的字符串

在程序的第7行中，mutate 函数接受 rune 切片作为参数。然后它将切片的第一个元素更改为'a'，
将 rune 转换成字符串并返回它。第14行中函数被调用。h 被转换为 rune切片并传递给 mutate。该程序输出 aello
*/
func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
