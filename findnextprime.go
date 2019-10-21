package piscine

func FindNextPrime(nb int) int {
	count := 0
	t := 0
	if nb == 0 || nb == 1 || nb < 0 {
		return 2
	}
	for i := 2; i < nb; i++ {
		t = nb % i
		if t == 0 {
			count++
			break
		}
	}
	if count >= 1 {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}
