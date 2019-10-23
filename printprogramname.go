package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	str := arguments[0]
	runes := []rune(str)
	for i := range runes {
		z01.PrintRune(runes[i])
	}

}
