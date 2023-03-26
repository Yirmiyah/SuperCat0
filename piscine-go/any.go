package piscine

func Any(f func(string) bool, a []string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			count++
		}
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}
