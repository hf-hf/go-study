package main

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

/**
Recover 仅在从同一 goroutine 调用时才起作用。
从不同的 goroutine 发生的 panic 中 Recover 是不可能的。让我们用一个例子来理解这一点。
*/
func main() {
	a()
	fmt.Println("normally returned from main")
}
