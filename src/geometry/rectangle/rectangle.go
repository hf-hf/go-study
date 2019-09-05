//rectprops.go
package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

func Area2(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal2(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}