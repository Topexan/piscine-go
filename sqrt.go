package piscine

func Sqrt(nb int) int {
	res := 1
	count := 0
	if nb == 0 || nb == 1 {
		return nb
	} else {
		for i := 1; res < nb; i++ {
			res = i * i
			count++
		}
		if res > nb {
			return 0
		} else {
			return count
		}
	}
}
