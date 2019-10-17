package piscine

func Swap(a *int, b *int) {
	pa := *a
	pb := *b

	*a = pb
	*b = pa
}