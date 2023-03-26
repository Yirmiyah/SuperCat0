package piscine

func IsSorted(f func(int, int) int, a []int) bool {
	croissant := true
	descendant := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			croissant = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			descendant = false
		}
	}

	return croissant || descendant
}

/*package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	for i := 1; i < len(a); i++ {

		if f(a[i-1], a[i]) == -1 {
			count++
		}
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}*/
