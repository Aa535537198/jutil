package jstr

import "strings"

// 判断多个字符串是否相等， 大小写敏感
func IsEqual(str1, str2 string, strs ...string) bool {
	return equal(strEqual, str1, str2, strs)
}

func strEqual(str1, str2 string) bool {
	return strings.Compare(str1, str2) == 0
}

//  判断多个字符串是否相等， 大小写不敏感
func IsEqualIgnoreCase(str1, str2 string, strs ...string) bool {
	return equal(strings.EqualFold, str1, str2, strs)
}

func equal(fun func(s1, s2 string) bool, str1, str2 string, strs []string) bool {
	if len(strs) == 0 {
		return fun(str1, str2)
	}
	f := true
	for i := 0; i < len(strs); i++ {
		f = fun(str1, strs[i])
		if !f {
			return f
		}
	}
	return str1 == str2 && f
}

// 判断字符串是否等于其中一个  大小写敏感
func IsEqualAny(str string, strs ...string) bool {
	return equalsAny(strEqual, str, strs)
}

// 判断字符串是否等于其中一个  大小写不敏感
func IsEqualAnyIgnoreCase(str string, strs ...string) bool {
	return equalsAny(strings.EqualFold, str, strs)
}

func equalsAny(fun func(string, string) bool, str string, strs []string) bool {
	if len(strs) == 0 {
		panic("input params is not enough")
	}
	f := false
	for i := 0; i < len(strs) && !f; i++ {
		f = fun(str, strs[i])
	}
	return f
}
