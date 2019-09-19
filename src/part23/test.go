package main

import (
	"fmt"
	"time"
)

var jobs = make(chan int, 2)

func test() {
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Println("jobs is close!")
}

func main() {
	go test()
	time.Sleep(3 * time.Second)
	for v := range jobs {
		fmt.Println(v)
	}
}
