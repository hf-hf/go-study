package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

/**
这是一个简单的程序，用于查找给定切片的最大数量。
largest 函数将 int 切片作为参数，并输出输入切片的最大数字。
largest 函数的第一行包含语句defer finished()。
这意味着在 largest 函数返回之前将调用 finished() 函数。
运行此程序，你可以看到以下输出。
*/
func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}
