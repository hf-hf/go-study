package main

import (
	"fmt"
)

func send(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func recevice(data chan int) {
	for {
		tmp, ok := <-data
		if ok == false {
			break
		}
		//time.Sleep(1 * time.Second)
		fmt.Printf("type: %t, value: %d, ok: %v \n", tmp, tmp, ok)
	}
}

func main() {
	data := make(chan int)
	go send(data)
	recevice(data)
}
