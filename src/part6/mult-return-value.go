package main

import (
	"fmt"
)

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	area, perimeter := rectProps2(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}