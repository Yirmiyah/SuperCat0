package piscine

func IsPrintable(s string) bool {
	tab := []rune(s)
	for i := 0; i < len(s); i++ {
		if tab[i] >= 32 && tab[i] <= 127 {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
