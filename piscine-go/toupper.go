package piscine

func ToUpper(s string) string {
	chaine := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			chaine[i] -= 32
		}
	}
	return string(chaine)
}
