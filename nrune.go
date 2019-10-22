package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	if n < 0 || n > len(word) {
		return ' '
	}
	return word[n-1]
}
