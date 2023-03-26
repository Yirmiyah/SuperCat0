package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nb int) int {
	if (nb % 2) == 0 {
		return 1
	} else {
		return 0
	}
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if !(isEven(lengthOfArg)) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
