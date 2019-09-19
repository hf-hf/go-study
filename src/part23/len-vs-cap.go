package main

import (
	"fmt"
)

/**
缓冲信道的容量是信道可以容纳的值的数量。这是我们使用 make 函数创建缓冲信道时指定的值。
缓冲信道的长度是当前在其中排队的元素个数。
*/
func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}
