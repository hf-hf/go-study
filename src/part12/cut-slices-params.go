package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

/**
这是一个语法糖，可用于将切片传递给可变参数函数。
你必须在切片后面使用 ... 。这样做的话，切片将直接传递给函数，而不会创建新切片。
*/
func main() {
	nums := []int{89, 90, 95}
	find(89, nums...)
}
