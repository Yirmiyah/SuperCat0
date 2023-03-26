package piscine

func IterativePower(nb int, power int) int {
	var puissance int
	puissance = 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		puissance = puissance * nb
	}
	return puissance
}
