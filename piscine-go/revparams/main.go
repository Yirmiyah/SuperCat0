package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}

func main() {
	var Arg string
	for i := len(os.Args) - 1; i >= 1; i-- {
		Arg = string(os.Args[i])
		PrintStr(Arg)
	}
}
