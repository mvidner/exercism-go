package collatzconjecture

import "errors"

// how many Collatz steps from *n* to one; error if *n* isn't positive
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be at least 1")
	}

	var steps int = 0
	for n != 1 {
		n = CollatzStep(n)
		steps += 1
	}
	return steps, nil
}

func CollatzStep(i int) int {
	if i%2 == 0 {
		return i / 2
	} else {
		return 3*i + 1
	}
}
