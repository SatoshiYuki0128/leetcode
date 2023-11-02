package easy

func MergeAlternately(word1 string, word2 string) string {
	result := ""
	index := 0

	size := smaller(len(word1), len(word2))

	for index < size {
		result += string(word1[index])
		result += string(word2[index])
		index++
	}

	if len(word1) == size {
		result += word2[index:]
	} else {
		result += word1[index:]
	}

	return result
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}
