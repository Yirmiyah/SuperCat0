package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	c := 0
	d := 0
	for index := range a {
		index = index
		c++
	}
	for index := range b {
		index = index
		d++
	}
	if d == 0 {
		return 0
	}

	for index, value := range a {
		if value == b[0] && c >= d+index-1 {
			run := 1
			for i := 1; i < d; i++ {
				if b[i] == a[index+i] {
					run++
				}
			}
			if run == d {
				return index
			}
		}
	}
	return -1
}

/*func Index(s string, toFind string) int {
	b := 0
	a := 0
	for ; a < len(s); a++ {
		if s[a] == toFind[b] {
			for ; b < len(toFind); b++ {
				if s[a] == toFind[b] && b == len(toFind)-1 {
					return a - b
				} else {
					a++
				}
			}
		}
	}
	return -1
}*/

/*func Index(s string, toFind string) int {
	j := 0
	i := 0
	count := 0

	for i = 0; i <= len(s)-1; i++ {
		for j = 0; j <= len(toFind)-1; j++ {
			if s[i] == toFind[j] {
				count++
				return i
			} else if count == 0 {
				return -1
			}
		}

	}
	return 0
}*/
