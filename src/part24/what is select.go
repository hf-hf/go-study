package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

/**
select 语句用于从多个发送/接收信道操作中进行选择。
select 语句将阻塞，直到其中一个发送/接收操作准备就绪。
如果准备好多个操作，则随机选择其中一个操作。
语法类似于 switch，但是每个 case 语句都是一个信道操作。让我们深入研究一些代码，以便更好地理解。

程序到达 select 语句。select 语句将阻塞，直到其中一个信道就绪。
在下面的程序中，server1 Goroutine 在6秒后写入 output1 信道，
而 server2 在3秒后写入 output2 信道。因此select语句将阻塞3秒，并等待 server2
 Goroutine写入output2 信道。3秒后，程序输出
然后程序终止。
*/
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
