package piscine

func BasicJoin(str []string) string {
	var sen string
	for i := range str {
		sen = sen + str[i]
	}
	return sen
}
