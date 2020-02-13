package series

// all the contiguous substrings of length *n* in *s*,
// in the order that they appear.
func All(n int, s string) []string {
	result := []string{}

	end := len(s) - n
	for start := 0; start <= end; start++ {
		result = append(result, UnsafeFirst(n, s[start:]))
	}
	return result
}

// the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// the first substring of s with length n, if possible
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
