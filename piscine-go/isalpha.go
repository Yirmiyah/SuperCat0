package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if ([]rune(s)[i]) <= 'Z' && ([]rune(s)[i]) >= 'A' || ([]rune(s)[i]) <= 'z' && ([]rune(s)[i]) >= 'a' || ([]rune(s)[i]) <= '9' && ([]rune(s)[i]) >= '0' {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

/*func IsAlpha(s string) bool {
	tab := []rune(s)
	//var count int
	for index := 0; index < len(s); index++ {
		if (tab[index] >= 97 && tab[index] <= 122) || (tab[index] >= 65 && tab[index] <= 90) || tab[index] >= 48 && tab[index] <= 57 {
			//count++
			if index == len(s)-1 {
				return true
			} else {
				return false
			}
		}

	}
	return false
}*/
