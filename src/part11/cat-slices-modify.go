package main

import (
	"fmt"
)

/**
在上面程序的第9行中，我们对数组的索引2,3,4创建了dslice 。
for循环将这些索引中的值递增1。当我们在for循环之后打印数组时，我们可以看到对切片的更改反映在数组中
*/
func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}
