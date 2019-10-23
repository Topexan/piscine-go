package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	var r int = 0
	for j := range runes {
		r++
		j = j
	}
	for i := 1; i < r; i++ {
		if runes[i-1] == ' ' || runes[i-1] == '+' {
			if runes[i] > 96 && runes[i] < 123 {
				runes[i] = runes[i] - 32
			}
		}
		i = i
	}
	str := string(runes)
	return str
}
