package main

import (
	"fmt"
)

/**
闭包是匿名函数的一种特殊情况。闭包是匿名函数，它访问函数体外部定义的变量。
*/
func main() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}
