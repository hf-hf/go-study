package main

import "fmt"

/**
每个信道都有一个与之关联的类型。这种类型是信道允许传输的数据类型。不允许使用信道传输其他类型。
chan T 是 T 类型的信道
信道的零值为 nil。nil 信道没有任何用处，因此必须使用 make 来定义信道，跟定义 maps 和 slices 类似。
*/
func main() {
	var a chan int
	//a := make(chan int)
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)

		/*
			箭头相对于信道的方向指定数据是发送还是接收。
			在第一行中，箭头从 a 指向内，因此我们从信道 a 读取数据并将值存储到变量数据。
			在第二行，箭头指向 a，因此我们向信道 a 写数据。

			发送和接收在默认情况下是阻塞的
			默认情况下，对信道的发送和接收是阻塞的。这是什么意思呢？当数据被发送到一个信道时，
			send 语句中被阻塞，直到其他 Goroutine 从该信道读取数据为止。
			类似地，当从信道读取数据时，读取将被阻塞，直到 Goroutine 将数据写入该信道。
			信道的这一特性有助于 goroutine 在不使用显式锁或条件变量的情况下有效地通信，
			而显式锁或条件变量在其他编程语言中非常常见。
		*/
		data := <-a // read from channel a
		a <- data   // write to channel a
	}
}
