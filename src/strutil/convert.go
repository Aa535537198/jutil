package strutil

import (
	"bytes"
	"unicode"
)

// 将字符串转变为字符切片
func ToCharSlice(str string) []string {
	bs := []byte(str)
	var res = make([]string, len(bs))

	for i := 0; i < len(bs); i++ {
		res[i] = string(bs[i])
	}
	return res
}

// 将 字符串切片转换为一个字符串
func ToString(strs ...string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(strs); i++ {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}

// 大小写转化
func ToSwapCase(str string) string {
	var buffer bytes.Buffer

	bs := []byte(str)

	for i := 0; i < len(bs); i++ {
		r := toRune(bs[i])
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				r = unicode.ToLower(r)
			} else if unicode.IsLower(r) {
				r = unicode.ToUpper(r)
			}
		}
		buffer.WriteString(string(r))
	}
	return buffer.String()
}
