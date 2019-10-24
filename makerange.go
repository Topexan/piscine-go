package piscine

func MakeRange(min, max int) []int {
	var ans []int
	l := 0
	if min < max {
		for i := min; i < max; i++ {
			l++
		}
		ans = make([]int, l)
		for j := 0; j < l; j++ {
			ans[j] = min
			min++
		}
	}
	return ans
}
