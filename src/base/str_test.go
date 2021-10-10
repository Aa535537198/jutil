package base

import (
	"fmt"
	"strings"
	"testing"
)

func TestJstr_Format(t *testing.T) {
	var s Jstr
	s = "hello world, my name id {} my school is {} "
	format := s.Format("{}", "shu")
	fmt.Println(format)
}

func TestJstr_ReplaceStringByIndex(t *testing.T) {
	var s Jstr
	s = "hello world, this is my testing welcome to..."

	index := strings.Index(string(s), "i")
	s = s.ReplaceStringByIndex("i", "z", index+1, -1, -1)
	fmt.Println(s)
}

func TestReplace(t *testing.T) {
	var s Jstr
	s = "hello world, this is my testing welcome to..."

	allString := s.ReplaceString("my", "jie")
	fmt.Println(allString)
}

func TestReplaceAllString(t *testing.T) {
	var s Jstr
	s = "hello world, this is my testing welcome to..."

	allString := s.ReplaceAllString(" jie", "my", "\\.\\.\\.")
	fmt.Println(allString)
}

func TestIsEmpty(t *testing.T) {
	var s Jstr
	s = "123332231321312assdadsadas"
	empty := s.isEmpty("[1-3]", "[a-z]")
	fmt.Println(empty)
}
