package main

import (
	"fmt"
)

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

/**
上面程序第17行中的函数调用将切片的每个元素减2。当
函数调用后输出切片时，这些更改是可见的。
如果你还记得话，这与数组不同，在数组中对函数内部的数组所做的更改在函数外部是不可见的
PS:固定长度的为数组，可变长度的为切片
arr：nos := [3]int{8, 7, 6}
cat-slices: nos := []int{8, 7, 6}
*/
func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
}
