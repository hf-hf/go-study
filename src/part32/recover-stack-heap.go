package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

/**
如果我们 recover 了 panic，我们就会丢失关于 panic 的堆栈跟踪。
即使在上述程序recover 后，我们也丢失了堆栈跟踪。
有一种方法可以使用 Debug 包的 PrintStack 函数打印堆栈跟踪。
*/
func main() {
	a()
	fmt.Println("normally returned from main")
}
