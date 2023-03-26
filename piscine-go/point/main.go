package main

import "github.com/01-edu/z01"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	res := "x = 42, y = 21"
	PrintStr(res)
	z01.PrintRune('\n')
}
