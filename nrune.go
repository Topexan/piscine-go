package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	var l int = 0
	for i := range word {
		l++
		i = i
	}
	if n < 0 || n > l {
		return ' '
	}
	return word[n-1]
}
