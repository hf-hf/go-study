package main

func sendData(sendch chan<- int) {
	sendch <- 10
}

/**
但是如果无法从发送信道读取数据，那么只向发送信道写入数据又有什么意义呢!
这就是使用信道转换的地方。可以将双向信道转换为仅发送或仅接收信道，但反之则不行。
*/
func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	//fmt.Println(<-sendch)
}
