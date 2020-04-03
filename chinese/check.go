package chinese

import "strings"

func IsChinese(s string) bool {
	return strings.Contains(s, "年") ||
		strings.Contains(s, "月") ||
		strings.Contains(s, "日") ||
		strings.Contains(s, "时") ||
		strings.Contains(s, "分") ||
		strings.Contains(s, "秒")
}

func IsWithoutTime(s string) bool {
	return !strings.Contains(s, " ")
}

func IsAbbrYear(s string) bool {
	return strings.Index(s, "年") == 2
}

func IsWithoutSecond(s string) bool {
	return !strings.Contains(s, "秒")
}
