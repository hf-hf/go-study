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
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}

/**
interface{} 的具体类型由 reflect.Type 表示，底层值由 reflect.Value 表示。
有两个函数 reflect.TypeOf() 和 reflect.ValueOf() 分别返回 reflect.Type 和 reflect.Value。
这两种类型是创建查询生成器的基础。让我们编写一个简单的例子来理解这两种类型。
*/
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}
