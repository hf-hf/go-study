package main

import (
	"fmt"
	"reflect"
)

/**
Int 和 String 方法有助于将 reflect.Value 分别提取为 int64 和 string 。
*/
func main() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}
