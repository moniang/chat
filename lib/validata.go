package lib

import (
	"regexp"
)

// 判断字段是否为空，其中有一个值不符合，则返回True
func IsEmpty(str ...string) bool {
	for _, s := range str {
		if len(s) == 0 {
			return true
		}
	}
	return false
}

// 判断字段是否由汉字、字母和数字组成，其中有一个值不符合，则返回False
func IsChsAlphaNum(str ...string) bool {
	for _, s := range str {
		if m, _ := regexp.MatchString("^[\u4e00-\u9fa5a-zA-Z0-9]+$", s); !m {
			return false
		}
	}
	return true
}

// 判断字段是否由字母和数字组成，其中有一个值不符合，则返回False
func IsAlphaNum(str ...string) bool {
	for _, s := range str {
		if m, _ := regexp.MatchString("^[A-Za-z0-9]+$", s); !m {
			return false
		}
	}
	return true
}

// 判断字段长度是否符合要求
func Len(min int, max int, str ...string) bool {
	for _, s := range str {
		if len(s) < min || len(s) > max {
			return false
		}
	}
	return true
}

// 判断字段是否为16进制的颜色值
func IsColorHex(str ...string) bool {
	for _, s := range str {
		if m, _ := regexp.MatchString("^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$", s); !m {
			return false
		}
	}
	return true
}
