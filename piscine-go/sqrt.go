package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	for i := 1; i < nb; i++ {
		if nb == i*i {
			if (nb % i) != 0 {
				return 0
			} else {
				return i
			}
		}
	}
	return 0
}
