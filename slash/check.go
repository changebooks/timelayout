package slash

import "strings"

func IsSlash(s string) bool {
	return strings.Contains(s, "/")
}

func IsWithoutTime(s string) bool {
	return !strings.Contains(s, " ")
}

func IsAbbrYear(s string) bool {
	return strings.Index(s, "/") == 2
}

func IsWithoutSecond(s string) bool {
	return strings.Count(s, ":") == 1
}
