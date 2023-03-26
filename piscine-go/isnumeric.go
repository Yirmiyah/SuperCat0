package piscine

func IsNumeric(s string) bool {
	for index := 0; index < len(s); index++ {
		if rune(s[index]) >= '0' && rune(s[index]) <= '9' {
			if index == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
