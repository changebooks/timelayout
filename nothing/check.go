package nothing

func IsNothing(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if '0' <= c && c <= '9' {
			continue
		}

		return false
	}

	return true
}

func IsWithoutTime(s string) bool {
	return len(s) <= 8
}
