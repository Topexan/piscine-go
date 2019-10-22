package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] > 64 && runes[i] < 91 {
			runes[i] = runes[i] + 32
		}
		i = i
	}
	str := string(runes)
	return str
}
