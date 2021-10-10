package strutil

import (
	"github.com/Aa535537198/jutil/src/convert"
	"regexp"
	"unicode"
)

// 去除掉一些字符后是否为空
func IsEmpty(this string, ignores ...string) bool {
	if len(this) != 0 && len(ignores) != 0 {
		for i := 0; i < len(ignores); i++ {
			compile := regexp.MustCompile(ignores[i])
			this = compile.ReplaceAllString(this, "")
		}
	}
	return len(this) == 0
}

func IsNotEmpty(this string, ignores ...string) bool {
	if len(this) != 0 && len(ignores) != 0 {
		for i := 0; i < len(ignores); i++ {
			compile := regexp.MustCompile(ignores[i])
			this = compile.ReplaceAllString(this, "")
		}
	}
	return len(this) != 0
}

func IsUpperCase(ch interface{}) bool {
	return unicode.IsUpper(toRune(ch))
}

// 判断是否为数字  能判断bytes数、半角与全角数字： unicode（=半角数字）
func IsDigit(ch interface{}) bool {
	return unicode.IsDigit(toRune(ch))
}

// 判断是否为字母
func IsLetter(ch interface{}) bool {
	return unicode.IsLetter(toRune(ch))
}

func toRune(ch interface{}) rune {
	var r rune
	switch ch.(type) {
	case int:
		r = rune(ch.(int))
	case string:
		r = rune(convert.ToInt(ch.(string)))
	case int32:
		r = ch.(int32)
	case int8:
		r = rune(ch.(int8))
	case int16:
		r = rune(ch.(int16))
	case int64:
		r = rune(ch.(int64))
	case byte:
		r = rune(ch.(byte))
	default:
		panic("toRune input param is error")
	}
	return r
}
