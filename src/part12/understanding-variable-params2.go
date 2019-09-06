package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	// 因为切片长度扩容时，是通过创建新的切片，然后拷贝到当前s上，因此新的元素不会追加到外部
	s = append(s, "playground")
	fmt.Println(s)
	fmt.Println(cap(s))
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	fmt.Println(cap(welcome))
}
