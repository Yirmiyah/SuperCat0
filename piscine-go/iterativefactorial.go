package piscine

func IterativeFactorial(nb int) int {
	if nb >= 99999 || nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	var factoriel int
	factoriel = 1
	for i := 1; i <= nb; i++ {
		factoriel = factoriel * i
	}
	return factoriel
}
