package textdistance

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	min := 0
	for i, v := range is {
		if i == 0 {
			min = v
		} else {
			if min > v {
				min = v
			}
		}
	}
	return min
}

// Max returns the maximum number of passed int slices.
func Max(is ...int) int {
	var max int
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
