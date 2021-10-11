package jstr

import (
	"regexp"
	"strings"
)

func Format(this, pattern string, params ...string) string {
	if len(params) < 1 {
		return this
	}
	for i := 0; i < len(params); i++ {
		this = ReplaceString(this, pattern, params[i], 1)
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

// 这两个是去掉字符串的前缀后缀的，例如去个文件名的扩展名啥。
func RemovePrefix(this, prefix string) string {
	hasPrefix := strings.HasPrefix(this, prefix)
	if hasPrefix {
		return this[len(prefix):]
	}
	return this
}

func RemoveSuffix(this, suffix string) string {
	hasSuffix := strings.HasSuffix(this, suffix)
	if hasSuffix {
		return this[:len(this)-len(suffix)]
	}
	return this
}

func RemovePrefixIgnoreCase(this, prefix string) string {
	tL := strings.ToLower(this)
	prefix = strings.ToLower(prefix)
	hasPrefix := strings.HasPrefix(tL, prefix)
	if hasPrefix {
		return this[len(prefix):]
	}
	return this
}

func RemoveSuffixIgnoreCase(this, suffix string) string {
	tL := strings.ToLower(this)
	suffix = strings.ToLower(suffix)
	hasSuffix := strings.HasSuffix(tL, suffix)
	if hasSuffix {
		return this[:len(this)-len(suffix)]
	}
	return this
}
