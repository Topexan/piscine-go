package piscine

func StrRev(s string) string {
	runes := []rune(s)
	var z rune
	var count int = 0
	for i := range runes {
		count++
		i = i
	}
	for l := 0; l < count/2; l++ {
		z = runes[l]
		runes[l] = runes[count-l-1]
		runes[count-l-1] = z
	}
	return string(runes)
}
