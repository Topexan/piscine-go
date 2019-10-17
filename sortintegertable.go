package piscine

func SortIntegerTable(table []int) {
	var length int = 0
	var z = 0
	for _, v := range table {
		v = v
		length++
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if table[i] > table[j] {
				z = table[i]
				table[i] = table[j]
				table[j] = z
			}
		}
	}
}
