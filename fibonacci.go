package piscine

func Fibonacci(power int) int {
	if power <= 0 {
		return 0
	} else if power == 1 {
		return 1
	} else {
		return Fibonacci(power-1) + Fibonacci(power-2)
	}
}
