package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		tab[i] = f(a[i])
	}
	return tab
}
