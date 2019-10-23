package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isFirst := true
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' && isFirst {
			runes[i] -= 32
			isFirst = false
		} else if (IsRuneDigit(runes[i]) || (runes[i] >= 'A' && runes[i] <= 'Z')) && isFirst {
			isFirst = false
		} else if !isFirst && runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if !isFirst && !IsRuneDigit(runes[i]) && !IsRuneAlpha(runes[i]) {
			isFirst = true
		}
	}
	return string(runes)
}
