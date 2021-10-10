package strutil

import (
	"regexp"
	"strings"
)

func Format(this string, params ...string) string {
	if len(params) < 2 {
		return this
	}
	old := params[0]
	for i := 1; i < len(params); i++ {
		this = ReplaceString(this, old, params[i], 1)
	}
	return this
}

// 替换字符串
func ReplaceAllString(this string, new string, replaces ...string) string {
	if len(replaces) == 0 {
		return this
	}

	for i := 0; i < len(replaces); i++ {
		compile := regexp.MustCompile(replaces[i])
		this = compile.ReplaceAllString(this, new)
	}

	return this
}

// 替换字符串
func ReplaceString(this string, old string, new string, counts ...int) string {
	if len(this) == 0 {
		return this
	}

	var c int
	if len(counts) == 0 {
		c = -1
	} else {
		c = counts[0]
	}
	return strings.Replace(this, old, new, c)
}

// 替换字符串依靠下标
func ReplaceStringByIndex(this string, old string, new string, startIndex, endIndex, count int) string {
	if len(this) == 0 {
		return this
	}
	prefix := this[:startIndex]
	if endIndex == -1 {
		this = this[startIndex:]
	} else {
		this = this[startIndex:endIndex]
	}

	return (prefix) + strings.Replace(this, old, new, count)
}

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
