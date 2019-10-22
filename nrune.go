package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	return word[n-1]
}
