package main

import (
	"fmt"
)

/**
当一个函数有多个延迟调用时，它们会被添加到堆栈中并以后进先出（LIFO）顺序执行。
*/
func main() {
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
