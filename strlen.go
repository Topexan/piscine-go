package piscine

func StrLen(str string) int {
	runes := []rune(str)
	var r int = 0
	for i := range runes {
		r++
		i = i
	}
	return r
}