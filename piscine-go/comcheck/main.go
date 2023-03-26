package main

import (
	"fmt"
	"os"
)

func main() {
	count := 0
	for i := 0; i < len(os.Args); i++ {
		argument := os.Args[i]
		if argument == "01" || argument == "galaxy 01" || argument == "galaxy" {
			count++
		}
	}
	if count > 0 {
		fmt.Println("Alert!!!")
	}
}
