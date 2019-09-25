package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func DecConvertToX(n, num int) (string, error) {
	if n < 0 {
		return strconv.Itoa(n), errors.New("只支持正整数")
	}
	if num != 2 && num != 8 && num != 16 {
		return strconv.Itoa(n), errors.New("只支持二、八、十六进制的转换")
	}
	result := ""
	h := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}
	for ; n > 0; n /= num {
		lsb := h[n%num]
		result = lsb + result
	}
	return result, nil
}

type FileMode2 uint32

func main() {
	fmt.Println(os.FileMode(0644))
	fmt.Println(os.ModeAppend.String())
	fmt.Println(1 << (32 - 1 - 2))
	fmt.Println(0777)
	fmt.Println(0b100000000)
	fmt.Println(int64(os.ModeType))
	fmt.Printf("Type is %T", os.ModeDir)
	// 0755->即用户具有读/写/执行权限，组用户和其它用户具有读写权限；
	// 0644->即用户具有读写权限，组用户和其它用户具有只读权限；
	fmt.Println(DecConvertToX(int(os.ModeDir), 8))
	fmt.Println(os.ModeDir)
	fmt.Println(os.FileMode(0644))
	// 0644 8进制的644 = 10进制的420
	v := uint32(0644)
	fmt.Println(v)
	tmp := os.FileMode(0644)
	fmt.Println(uint32(tmp.Perm()))
	// os.OpenFile(文件名，打开方式，打开模式）
	// 追加和只写模式
	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
