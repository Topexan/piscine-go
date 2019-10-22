package piscine

func LastRune(s string) rune {
	word := []rune(s)
	var l int = 0
	for i := range word {
		l++
		i = i
	}
	return word[l-1]
}
