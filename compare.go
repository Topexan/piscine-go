package piscine

func Compare(a, b string) int {
	runes1 := []rune(a)
	runes2 := []rune(b)
	var l1 int = 0
	var l2 int = 0
	for i := range runes1 {
		l1++
		i = i
	}
	for j := range runes2 {
		l2++
		j = j
	}
	num := l1
	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		return -1
	} else if l2 == 0 {
		return 1
	} else {
		if l1 > l2 {
			num = l2
		}
		for k := 0; k < num; k++ {
			if runes1[k] > runes2[k] {
				return 1
			} else if runes1[k] < runes2[k] {
				return -1
			}
		}
		if l1 == l2 {
			return 0
		} else if l1 > l2 {
			return 1
		}
		return -1
	}
}
