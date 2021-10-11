package jstr

import (
	"fmt"
	"github.com/Aa535537198/jutil/src/convert"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	s := "hello   World {}"
	prefix := RemovePrefix(s, "hello")
	fmt.Println(prefix)

	suffix := RemoveSuffixIgnoreCase(s, "world")
	fmt.Println(suffix)

	prefix = RemovePrefix(s, "hello   worl")
	fmt.Println(prefix)
	i := IsEmpty(prefix)
	fmt.Println(i)

	format := Format(s, "{}", "jie")
	fmt.Println(format)
}

func TestIsEqual(t *testing.T) {
	s1 := "hello"
	s2 := "hello"
	s3 := "hellO"

	equal := IsEqual(s1, s2, s3)
	fmt.Println(equal)

	ignoreCase := IsEqualIgnoreCase(s1, s2, s3)
	fmt.Println(ignoreCase)
}

func TestIsEqualAny(t *testing.T) {
	s1 := "hello"
	s2 := "hellO"
	s3 := "hellO2"

	strs := []string{"234", "2"}

	equalAny := IsEqualAny("234", strs...)
	fmt.Println(equalAny)

	any := IsEqualAny(s1, s2, s3)
	fmt.Println(any)

	any = IsEqualAnyIgnoreCase(s1, s2, s3)
	fmt.Println(any)
}

func TestToCharSlice(t *testing.T) {
	s := "hello World aDI sHUANG jIE"

	slice := ToCharSlice(s)
	fmt.Println(slice)

	ss := []byte(s)

	fmt.Println(ss)

	toString := ToString(slice...)
	fmt.Println(toString)

	swapCase := ToSwapCase(s)
	fmt.Println(swapCase)
}

func TestFormat(t *testing.T) {
	i := 10
	str := convert.ToStr(i)
	fmt.Println(str)

	i = 97

	k := string(rune(i))
	fmt.Println(k)

	fmt.Println()
}
