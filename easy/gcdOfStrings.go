package easy

func GcdOfStrings(str1 string, str2 string) string {
	result := ""

	if str1+str2 != str2+str1 {
		return result
	}

	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
