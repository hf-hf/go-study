package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

/**
WaitGroup 是一种结构类型，我们在第18行创建一个 WaitGroup 类型的零值变量。
WaitGroup 的工作方式是使用计数器。当我们在 WaitGroup 上调用 Add 并将其传递给 int 时，WaitGroup 的计数器会被传递给 Add 的值递增。递减计数器的方法是在WaitGroup 上调用 Done() 方法。Wait() 方法阻塞调用它的 Goroutine ，直到计数器变为零。
在上面的程序中，我们在第20行中调用 wg.Add(1)。for 循环中迭代3次。所以计数器现在变成3。for 循环也产生3个 process Goroutines，然后在第23行中调用wg.Wait()。 main Goroutine 等到计数器变为零。在第13行的 process Goroutine中，通过调用 wg.Done() 来减少计数器。 一旦所有3个派生的 Goroutines 完成执行，那就是 wg.Done() 被调用三次，计数器将变为零，main Goroutine 将被解锁。
在第27行传递wg的地址是很重要的。如果没有传递地址，那么每个Goroutine将拥有自己的 WaitGroup  副本，并且当它们执行完毕时，main将不会收到通知。
*/
func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		// 目的为使用同一个waitGroup
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
