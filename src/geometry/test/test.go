//rectprops.go
package test

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

func Area3(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal3(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}