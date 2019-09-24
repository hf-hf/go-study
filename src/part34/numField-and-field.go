package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

/**
在上面的程序中，我们第14行首先检查 q 的 Kind 是否是 struct ，
因为 NumField 方法仅适用于 struct 。 剩下的程序一目了然。
*/

/**
NumField() 方法返回结构中的字段数，Field(i int) 方法返回第 i 个字段的reflect.Value。
*/
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}
