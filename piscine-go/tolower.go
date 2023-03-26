package piscine

func ToLower(s string) string {
	chaine := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			chaine[i] += 32
		}
	}
	return string(chaine)
}
