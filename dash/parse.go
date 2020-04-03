package dash

import "time"

// 字符串 => time.Time
func Parse(s string) (time.Time, error) {
	return time.Parse(ParseLayout(s), s)
}

// 分析条件，获取模板
func ParseLayout(s string) string {
	return ParseLayoutEx(IsWithoutTime(s), IsAbbrYear(s), IsWithoutSecond(s))
}
