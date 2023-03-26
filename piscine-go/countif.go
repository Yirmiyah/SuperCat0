package piscine

func CountIf(f func(string) bool, a []string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			count++
		}
	}
	return count
}
