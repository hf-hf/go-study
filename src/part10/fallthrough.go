package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

/**
	Switch和case表达式不一定只是常量。它们也可以在运行时进行判断。
	在上面的程序中，num 初始化为函数 number() 的返回值。
	控件位于 switch 内部，并对 case 进行判断。case num < 100: 为真，
	程序输出 75 is lesser than 100。下一个语句是fallthrough。
	当遇到 fallthrough 时，控件将移动到下一个 case 的第一个语句，
	并输出75 is lesser than 200。程序的输出为
 */
func main() {
	switch num := number(); { //num is not a cofallthrough 应该在一个case中最后声明， 如果它出现在中间的某个地方，编译器将抛出错误 fallthrough statement out of placenstant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}