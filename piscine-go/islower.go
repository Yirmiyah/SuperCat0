package piscine

func IsLower(s string) bool {
	tab := []rune(s)
	count := 0
	for i := 0; i <= len(s)-1; i++ {
		if tab[i] >= 97 && tab[i] <= 122 || tab[i] >= 48 && tab[i] <= 57 || tab[i] >= 65 && tab[i] <= 90 {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
