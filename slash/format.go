package slash

import "time"

// 当前时间 => 字符串，条件：包含时间？简写年？包含秒？
func Now(withoutTime bool, abbrYear bool, withoutSecond bool) string {
	return FormatEx(time.Now(), withoutTime, abbrYear, withoutSecond)
}

// 时间戳 => 字符串，条件：包含时间？简写年？包含秒？
func Format(second int64, withoutTime bool, abbrYear bool, withoutSecond bool) string {
	return FormatEx(time.Unix(second, 0), withoutTime, abbrYear, withoutSecond)
}

// time.Time => 字符串，条件：包含时间？简写年？包含秒？
func FormatEx(t time.Time, withoutTime bool, abbrYear bool, withoutSecond bool) string {
	return t.Format(FormatLayoutEx(withoutTime, abbrYear, withoutSecond))
}
