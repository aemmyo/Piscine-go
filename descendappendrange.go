package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{} // return a non-nil empty slice
	}

	result := []int{} // initialize as empty slice, not nil
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
