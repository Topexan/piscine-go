package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	var result int = 0
	isnum := 0
	isneg := 0
	runes := []rune(s)
	for _, j := range runes {
		if j > 47 && j < 58 || j == '-' {
			result *= 10
			if j == '0' {
				result += 0
			} else if j == '1' {
				result += 1
				isnum++
			} else if j == '2' {
				result += 2
				isnum++
			} else if j == '3' {
				result += 3
				isnum++
			} else if j == '4' {
				result += 4
				isnum++
			} else if j == '5' {
				result += 5
				isnum++
			} else if j == '6' {
				result += 6
				isnum++
			} else if j == '7' {
				result += 7
				isnum++
			} else if j == '8' {
				result += 8
				isnum++
			} else if j == '9' {
				result += 9
				isnum++
			} else if j == '-' && isnum == 0 {
				isneg++
			} else if j == '-' {
				result /= 10
			}
		}
	}
	if isneg > 0 {
		result *= -1
	}
	return result
}
