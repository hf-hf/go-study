package main

import (
	"fmt"
)

/**
Go还提供了一个方便的函数 new 来创建指针。new 函数接受一个type作为参数，
返回一个新指针（值为这个type类型的零值）。
*/
func main() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}
