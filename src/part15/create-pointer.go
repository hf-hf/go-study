package main

import (
	"fmt"
)

/**
*T 是指针变量的类型，它指向类型 T 的值。
&运算符用于获取变量的地址。在上面的程序第9行中，我们将b的地址赋值给类型为*int的a。
现在a指向b，当我们输出a的值时，会打印出b的地址。
*/
func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
