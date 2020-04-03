package chinese

const (
	LayoutFormat                      = "2006年01月02日 15时04分05秒"
	LayoutFormatWithoutTime           = "2006年01月02日"
	LayoutFormatAbbrYear              = "06年01月02日 15时04分05秒"
	LayoutFormatWithoutSecond         = "2006年01月02日 15时04分"
	LayoutFormatAbbrYearWithoutSecond = "06年01月02日 15时04分"
	LayoutFormatWithoutTimeAbbrYear   = "06年01月02日"
)

const (
	LayoutParse                      = "2006年1月2日 15时4分5秒"
	LayoutParseWithoutTime           = "2006年1月2日"
	LayoutParseAbbrYear              = "06年1月2日 15时4分5秒"
	LayoutParseWithoutSecond         = "2006年1月2日 15时4分"
	LayoutParseAbbrYearWithoutSecond = "06年1月2日 15时4分"
	LayoutParseWithoutTimeAbbrYear   = "06年1月2日"
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
