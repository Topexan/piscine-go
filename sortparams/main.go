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
	runes := make([]rune, l-1)
	for j := 1; j < l; j++ {
		str := []rune(arguments[j])
		runes[j-1] = str[0]
	}

	r := ' '
	for k := 0; k < l; k++ {
		for z := k + 1; z < l-1; z++ {
			if runes[k] > runes[z] {
				r = runes[k]
				runes[k] = runes[z]
				runes[z] = r
			}
		}
	}
	for x := range runes {
		z01.PrintRune(runes[x])
		z01.PrintRune(10)
	}

	//	z01.PrintRune(rune(l))
}
