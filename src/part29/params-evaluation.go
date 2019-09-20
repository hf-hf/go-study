package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

/**
延迟函数的参数在执行defer语句时计算，而不是在实际函数调用完成时计算。
*/
func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}
