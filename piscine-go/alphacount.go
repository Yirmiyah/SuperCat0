package piscine

func AlphaCount(s string) int {
	array := []rune(s)
	n := 0
	for i := 0; i <= len(s)-1; i++ {
		if array[i] >= 65 && array[i] <= 90 || array[i] >= 97 && array[i] <= 122 {
			n++
		}
	}
	return n
}
