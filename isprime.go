package piscine

func IsPrime(nb int) bool {
	count := 0
	t := 0
	if nb == 0 || nb == 1 || nb < 0 {
		return false
	}
	for i := nb - 1; i > 1; i-- {
		t = nb % i
		if t == 0 {
			count++
		}
	}
	if count >= 1 {
		return false
	} else {
		return true
	}
}
