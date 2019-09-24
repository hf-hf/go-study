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
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
}

/**
Type 表示 interface{} 的实际类型，
在这种情况下，main.Order 和 Kind 表示类型的特定种类。在这种情况下，它是一个 struct。
*/
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}
