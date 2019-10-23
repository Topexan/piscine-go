package piscine

func Join(str []string, sep string) string {
	var sen string
	r := 0
	for i := range str {
		r++
		i = i
	}
	for j := 0; j < r-1; j++ {
		sen = sen + str[j] + sep
	}
	return sen + str[r-1]
}
