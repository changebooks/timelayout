package dash

const (
	LayoutFormat                      = "2006-01-02 15:04:05"
	LayoutFormatWithoutTime           = "2006-01-02"
	LayoutFormatAbbrYear              = "06-01-02 15:04:05"
	LayoutFormatWithoutSecond         = "2006-01-02 15:04"
	LayoutFormatAbbrYearWithoutSecond = "06-01-02 15:04"
	LayoutFormatWithoutTimeAbbrYear   = "06-01-02"
)

const (
	LayoutParse                      = "2006-1-2 15:4:5"
	LayoutParseWithoutTime           = "2006-1-2"
	LayoutParseAbbrYear              = "06-1-2 15:4:5"
	LayoutParseWithoutSecond         = "2006-1-2 15:4"
	LayoutParseAbbrYearWithoutSecond = "06-1-2 15:4"
	LayoutParseWithoutTimeAbbrYear   = "06-1-2"
)

// 选择Format模板，条件：包含时间？简写年？包含秒？
func FormatLayoutEx(withoutTime bool, abbrYear bool, withoutSecond bool) string {
	if withoutTime {
		if abbrYear {
			return LayoutFormatWithoutTimeAbbrYear
		} else {
			return LayoutFormatWithoutTime
		}
	} else {
		if abbrYear {
			if withoutSecond {
				return LayoutFormatAbbrYearWithoutSecond
			} else {
				return LayoutFormatAbbrYear
			}
		} else {
			if withoutSecond {
				return LayoutFormatWithoutSecond
			} else {
				return LayoutFormat
			}
		}
	}
}

// 选择Parse模板，条件：包含时间？简写年？包含秒？
func ParseLayoutEx(withoutTime bool, abbrYear bool, withoutSecond bool) string {
	if withoutTime {
		if abbrYear {
			return LayoutParseWithoutTimeAbbrYear
		} else {
			return LayoutParseWithoutTime
		}
	} else {
		if abbrYear {
			if withoutSecond {
				return LayoutParseAbbrYearWithoutSecond
			} else {
				return LayoutParseAbbrYear
			}
		} else {
			if withoutSecond {
				return LayoutParseWithoutSecond
			} else {
				return LayoutParse
			}
		}
	}
}
