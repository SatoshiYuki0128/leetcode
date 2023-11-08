package easy

// -5, 1, 5, 0, -7

func LargestAltitude(gain []int) int {
	alt := 0
	result := alt

	for _, v := range gain {
		alt += v

		//fmt.Println(alt)

		if alt > result {
			result = alt
		}
	}

	return result
}
