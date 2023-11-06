package easy

func IsSubsequence(s string, t string) bool {
	sizeS, sizeT := len(s), len(t)

	if sizeS == 0 {
		return true
	}

	if sizeT == 0 {
		return false
	}

	if sizeT < sizeS {
		return false
	}

	indexS, indexT := 0, 0

	for indexT < sizeT {
		if s[indexS] == t[indexT] {
			indexS++
		}

		indexT++

		if indexS >= sizeS {
			return true
		}

	}

	return false
}
