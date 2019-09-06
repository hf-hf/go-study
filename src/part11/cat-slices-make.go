package main

import (
	"fmt"
)

/**
func make([]T, len, cap) []T 可以通过传递类型，长度和容量来创建切片。
capacity参数是可选的，默认为length。 make函数创建一个数组并返回一个切片引用。
*/
func main() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}
