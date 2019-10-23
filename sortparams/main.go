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
	r := " "
	for k := 1; k < l; k++ {
		for z := k + 1; z < l; z++ {
			str1 := []rune(arguments[k])
			str2 := []rune(arguments[z])
			rune1 := str1[0]
			rune2 := str2[0]
			if rune1 > rune2 {
				r = arguments[k]
				arguments[k] = arguments[z]
				arguments[z] = r
			}
		}
	}
	for j := 1; j < l; j++ {
		strfin := []rune(arguments[j])
		runefin := strfin[0]
		z01.PrintRune(runefin)
		z01.PrintRune(10)
	}
}
