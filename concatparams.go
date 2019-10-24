package piscine

func ConcatParams(args []string) string {
	var str string
	l := 0
	for i := range args {
		l++
		i = i
	}
	for j := 0; j < l-1; j++ {
		str = str + args[j] + "\n"
	}
	return str + args[l-1]
}
