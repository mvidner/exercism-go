package hamming

import "errors"

func Distance(a, b string) (int, error) {
	difference_count := 0
	alen := len(a)
	blen := len(b)
	if alen != blen {
		return 0, errors.New("Distance only works on strings of equal length")
	}
	for i := 0; i < alen; i++ {
		if a[i] != b[i] {
			difference_count++
		}
	}
	return difference_count, nil
}
