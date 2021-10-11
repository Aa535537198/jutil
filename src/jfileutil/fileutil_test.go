package jfile

import (
	"fmt"
	"testing"
)

func TestReadFileByOnce(t *testing.T) {
	path := "D:\\Users\\53553\\go\\src\\jutil\\jutil\\README.md"
	//s, _ := ReadFileByOnce(path)
	//fmt.Println(s)

	lines := ReadFileByLines(path, 3, 4)

	fmt.Println(lines)
	for _, val := range lines {
		fmt.Println(val)
	}

	fmt.Println("**********************")

	i, _ := ReadFileLinesByPath(path)
	fmt.Println(i)
}
