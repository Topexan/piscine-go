package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, str := range table {
		PrintStr(str)
		z01.PrintRune('\n')
	}
}
