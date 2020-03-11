package lib

// 判断字段是否为空，其中有一个值不符合，则返回True
func IsEmpty(str ...string) bool {
	for _, s := range str {
		if len(s) == 0 {
			return true
		}
	}
	return false
}
