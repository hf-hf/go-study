package main

import "fmt"

func main() {
	// 若信道未被初始化，使用select无default会报错死锁
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}
}
