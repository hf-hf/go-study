package main

import "fmt"

func multFunc(value int) int {
	return value * 5
}

func iterFunc(a *[]int, f func(int) int) {
	for index, value := range *a {
		(*a)[index] = f(value)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	iterFunc(&a, multFunc)
	fmt.Println(a)
}
