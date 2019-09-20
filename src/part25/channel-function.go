package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

/**
互斥锁 vs 信道
我们同时使用互斥锁和信道解决了竞争条件问题。那么我们如何决定什么时候使用什么。
答案在于你要解决的问题。如果试图解决的问题更适合互斥锁，那么继续使用互斥锁。如果真的需要，不要犹豫使用互斥锁。如果这个问题似乎更适合信道来解决，那么也可以使用它:)。
大多数Go新手尝试使用信道来解决每个并发问题，因为信道是该语言的一个很酷的特性。
这是错误的。该语言让我们可以选择使用互斥锁或信道，选择这两种方式都没有错。
一般来说，当 Goroutine 需要彼此通信时使用信道，当只有一个Goroutine应该访问代码的临界区时使用互斥锁。
在我们上面解决的问题中，我宁愿使用互斥锁，因为这个问题不需要 goroutine 之间的任何通信。
因此互斥锁是一个适合的选择。
我的建议是，为问题选择工具，不要试图将问题适应于工具:)
*/
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
