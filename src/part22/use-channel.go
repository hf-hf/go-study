package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

/**
代码第13行我们创建了一个 done bool信道。
并将其作为参数传递给 hello Goroutine。我们正在从 done 信道接收数据。
这行代码是阻塞的，这意味着在Goroutine将数据写入 done 信道之前，程序不会执行下一行代码。
因此，我们就不需要原始程序中的 time.Sleep 了。
代码行 <-done 从done 信道接收数据，但不在任何变量中使用或存储该数据。
这是完全合法的。
现在我们的 main Goroutine 被阻塞等待 done 信道的数据。
hello  Goroutine接收此信道作为参数，输出Hello world
goroutine然后写入 done 信道。当这个写入完成时，main Goroutine 从 done 信道接收数据，
它被解锁，然后输出 main function。
*/
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
