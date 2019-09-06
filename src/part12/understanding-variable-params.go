package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
}

/**
在程序第13行中，我们使用语法糖 ... 并将切片作为变量参数传递给 change 函数。
正如我们已经讨论的那样，如果使用了... ，welcome 切片本身将作为参数传递，而不会创建新的切片。
因此welcome将作为参数传递给 change 函数。
在change函数内部，切片的第一个元素更改为Go。因此该程序输出[Go world]
*/
func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
