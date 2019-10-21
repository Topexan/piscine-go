package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb <= 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			res = res * i
		}
		return res
	}
}
