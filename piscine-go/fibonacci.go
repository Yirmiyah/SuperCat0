package piscine

func Fibonacci(index int) int {
	a := 0
	b := 1
	if index == 0 {
		return 0
	}
	if index < 0 {
		return -1
	}
	if index > 1 {
		for i := 0; i < index; i++ {
			temp := a
			a = b
			b = temp + a
		}
	}
	return a
}
