package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	nameFile := os.Args[1]
	if len(os.Args) > 1 {
		a, _ := ioutil.ReadFile(nameFile)
		fmt.Print(string(a))
	} else {
		fmt.Print("Too many arguments")
	}
}
