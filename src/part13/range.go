package main

import (
	"fmt"
)

/**
一个要注意的的点是，在使用 for range 时，不能保证每次执行程序时，字典检索值的顺序都是相同的。
*/
func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}
