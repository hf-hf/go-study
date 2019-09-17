package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

/**
发送方能够关闭信道，通知接收方信道上不再发送数据。
接收方可以在接收来自信道的数据时使用附加变量来检查信道是否已关闭。
在上面的语句中，如果通过成功的发送操作接收到某个信道的值，ok为true。
如果ok为false，则表示我们从一个已经关闭的信道读取值。
从关闭的信道读取的值是信道类型的零值。例如，如果信道是 int 类型的信道，那么从关闭的新道接收的值将为 0。
*/
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
