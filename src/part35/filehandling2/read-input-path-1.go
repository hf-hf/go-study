package main

import (
	"flag"
	"fmt"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	// 应该在程序访问任何标志之前调用 flag.Parse()
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}
