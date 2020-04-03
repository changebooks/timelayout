package nothing

import "time"

// 当前时间 => 字符串，条件：包含时间？
func Now(withoutTime bool) string {
	return FormatEx(time.Now(), withoutTime)
}

// 时间戳 => 字符串，条件：包含时间？
func Format(second int64, withoutTime bool) string {
	return FormatEx(time.Unix(second, 0), withoutTime)
}

// time.Time => 字符串，条件：包含时间？
func FormatEx(t time.Time, withoutTime bool) string {
	return t.Format(FormatLayoutEx(withoutTime))
}
