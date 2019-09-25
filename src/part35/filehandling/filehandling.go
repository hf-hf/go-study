package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("../filehandling")
	data := box.String("test.txt")
	fmt.Println("Contents of file:", data)
}
