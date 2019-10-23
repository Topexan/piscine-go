package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	l := 0
	for i := range arguments {
		l++
		i = i
	}
	for j := l - 1; j > 0; j-- {
		runes := []rune(arguments[j])
		for k := range runes {
			z01.PrintRune(runes[k])
		}
		z01.PrintRune(10)
	}
}
