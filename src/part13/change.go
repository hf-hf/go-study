package main

import (
	"fmt"
)

/**
与切片类似，字典也是引用类型。
当将字典分配给一个新变量时，它们都指向相同的内部数据结构。因此，一个字典改变了值也会影响另外一个字典。
*/
func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
}
