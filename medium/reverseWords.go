package medium

// ReverseWords leetcode 151
func ReverseWords(s string) string {
	result := ""

	if len(s) == 0 {
		return result
	}

	var array []string
	index := 0
	end := 0

	// "  hello world  "

	for index < len(s) {
		if s[index] == ' ' {
			index++
			end = index
		} else {
			if s[end] == ' ' {
				array = append(array, s[index:end])
				index = end
			} else {
				end++
			}
		}

		if end == len(s) {
			if s[index:end] != "" {
				array = append(array, s[index:end])
			}
			break
		}

	}

	//fmt.Println(array)

	for i := len(array) - 1; i >= 0; i-- {
		result += " " + array[i]
	}

	return result[1:]
}
