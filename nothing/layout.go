package nothing

const (
	LayoutFormat            = "20060102150405"
	LayoutFormatWithoutTime = "20060102"
)

const (
	LayoutParse            = "20060102150405"
	LayoutParseWithoutTime = "20060102"
)

// 选择Format模板，条件：包含时间？
func FormatLayoutEx(withoutTime bool) string {
	if withoutTime {
		return LayoutFormatWithoutTime
	} else {
		return LayoutFormat
	}
}

// 选择Parse模板，条件：包含时间？
func ParseLayoutEx(withoutTime bool) string {
	if withoutTime {
		return LayoutParseWithoutTime
	} else {
		return LayoutParse
	}
}
