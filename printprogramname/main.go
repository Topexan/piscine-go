package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	str := arguments[0]
	name := []rune(str)
	for i := range name {
		z01.PrintRune(name[i])
	}
	z01.PrintRune(10)
}
