package main

import (
	"fmt"
)

/**
在上面的程序中，fruitslice 是从 fruitarray 索引1和2切出来的切片。
所以fruitslice 长度为2。fruitarray 的长度是7。fruiteslice 是从fruitarray 的索引1开始切的。
因此，fruiteslice 的容量是指从索引1 orange 开始的。
因此，fruiteslice 的容量为6。程序输出长度为2，容量为6的切片。
*/
func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
}
