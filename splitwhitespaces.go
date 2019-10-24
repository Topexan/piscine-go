package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str)
	spaces := 1
	lrunes := 0
	for i := range runes {
		lrunes++
		if runes[i] == ' ' {
			spaces++
		}
	}
	var words = make([]string, spaces)
	x := 0
	for j := 0; j < spaces; j++ {
		for k := x; k < lrunes; k++ {
			if runes[k] == ' ' {
				x++
				break
			}
			x++
			words[j] = words[j] + string(runes[k])
		}
	}
	return words
}
