package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		rt := reflect.TypeOf(q)
		t := rt.Name()
		v := reflect.ValueOf(q)
		fieldNum := v.NumField()
		colums := make([]string, 0, fieldNum)
		fmt.Println("slice大小:", len(colums))
		fmt.Println("slice容量:", cap(colums))
		for i := 0; i < fieldNum; i++ {
			colums = append(colums, rt.Field(i).Name)
		}
		fmt.Println("slice大小:", len(colums))
		fmt.Println("slice容量:", cap(colums))
		t += "(" + strings.Join(colums, ",") + ")"
		query := fmt.Sprintf("insert into %s values(", t)
		for i := 0; i < fieldNum; i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

/**
是否应该使用反射?
展示了反射的实际用途，现在才是真正的问题。 你应该使用反射吗？
我想引用Rob Pike关于使用反射的格言来回答这个问题。
Clear is better than clever.
Reflection is never clear.
在Go中，反射是一个非常强大和先进的概念，应该谨慎使用。
使用反射编写清晰且可维护的代码非常困难。 应尽可能避免使用，并且只有在绝对必要时才应使用反射。
*/
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	//_ = GetFieldName(o)
	createQuery(o)
	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}
