package medium

import (
	"strconv"
)

func Compress(chars []byte) int {
	size := len(chars)

	if size < 2 {
		return size
	}

	result := ""

	index := 1
	count := 1

	for index < size {
		if chars[index] == chars[index-1] {
			count++
		} else {
			result += string(chars[index-1])
			if count > 1 {
				result += strconv.Itoa(count)
			}
			count = 1
		}

		index++

		if index == size {
			result += string(chars[index-1])

			if count > 1 {
				result += strconv.Itoa(count)
			}
		}
	}

	index = 0

	for index < len(result) {
		chars[index] = result[index]
		index++
	}

	chars = chars[:index]
	return len(chars)
}
