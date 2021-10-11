package jprecondition

// 判断是否相等
func IsEqual(flag bool, msg interface{}) {
	if flag {
		return
	}
	panic(msg)
}
