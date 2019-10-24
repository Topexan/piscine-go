package piscine

func AppendRange(min, max int) []int {
	var ans []int
	if min < max {
		for i := min; i < max; i++ {
			ans = append(ans, i)
		}
	}
	return ans
}
