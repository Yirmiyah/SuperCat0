package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myProgramName := os.Args[0]
	array := []rune(myProgramName)
	affiche := []rune{}

	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == 47 {
			i = -1
		} else {
			affiche = append(affiche, array[i])
		}
	}
	for i := len(affiche) - 1; i >= 0; i-- {
		z01.PrintRune(rune(affiche[i]))
	}
	z01.PrintRune('\n')
}
