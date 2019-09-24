package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

/**
在上面的程序中，函数 appendStr 返回一个闭包。该闭包绑定到变量 t。让我们明白这意味着什么。
17 18行中声明的变量 a 和 b 是闭包，它们与自身的 t 值绑定。
我们先用参数 World 调用 a。 现在，一个版本的 t 的值变成了 Hello World。
我们在20行传进参数 Everyone 调用 b 函数。 由于 b 绑定到自己的变量 t，
因此 b的 t 版本的初始值为 Hello 。 因此，在此函数调用之后，b 的t版本的值变为 Hello Everyone。
剩下的程序就一目了然了。
*/
