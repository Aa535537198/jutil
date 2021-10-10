package base

import (
	"regexp"
	"strings"
)

type Jstr string

func (this Jstr) Format(params ...string) Jstr {
	if len(params) < 2 {
		return this
	}

	old := params[0]
	for i := 1; i < len(params); i++ {
		this = this.ReplaceString(old, params[i], 1)
	}
	return this
}

// 替换字符串
func (this Jstr) ReplaceAllString(new string, replaces ...string) Jstr {
	if len(replaces) == 0 {
		return this
	}

	for i := 0; i < len(replaces); i++ {
		compile := regexp.MustCompile(replaces[i])
		this = Jstr(compile.ReplaceAllString(string(this), new))
	}
	return this
}

// 替换字符串
func (this Jstr) ReplaceString(old string, new string, counts ...int) Jstr {
	if this.length() == 0 {
		return this
	}

	var c int
	if len(counts) == 0 {
		c = -1
	} else {
		c = counts[0]
	}
	return Jstr(strings.Replace(string(this), old, new, c))
}

// 替换字符串依靠下标
func (this Jstr) ReplaceStringByIndex(old string, new string, startIndex, endIndex, count int) Jstr {
	if this.length() == 0 {
		return this
	}
	prefix := this[:startIndex]
	if endIndex == -1 {
		this = this[startIndex:]
	} else {
		this = this[startIndex:endIndex]
	}

	return Jstr(string(prefix) + strings.Replace(string(this), old, new, count))
}

// 去除掉一些字符后是否为空
func (this Jstr) isEmpty(ignores ...string) bool {
	if this.length() != 0 && len(ignores) != 0 {
		for i := 0; i < len(ignores); i++ {
			compile := regexp.MustCompile(ignores[i])
			this = Jstr(compile.ReplaceAllString(string(this), ""))
		}
	}
	return this.length() == 0
}

// 获得Jstr的长度
func (this *Jstr) length() int {
	return len(*this)
}
