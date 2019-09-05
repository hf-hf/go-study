package main

import "fmt"

/**
Go中的数组是值类型而不是引用类型。这意味着当将它们分配给新变量时，会将原始数组的副本分配给新变量。
如果对新变量进行了更改，则它不会在原始数组中变化。
*/
func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
