package easy

import "strings"

func ReverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	array := []rune(s)

	left, right := 0, len(array)-1

	for left < right {
		if !strings.Contains(vowels, string(array[left])) {
			left++
			continue
		}

		if !strings.Contains(vowels, string(array[right])) {
			right--
			continue
		}

		array[left], array[right] = array[right], array[left]
		left++
		right--
	}

	return string(array)
}
