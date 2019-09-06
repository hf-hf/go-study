package main

/**
不能使用==运算符来比较字典。== 只能用于检查字典是否为 nil。
检查两个字典是否相等的一种方法是逐个比较每个字典的元素。我建议你可以写一个程序来验证它们:)。
*/
func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	if map1 == map2 {
	}
}
