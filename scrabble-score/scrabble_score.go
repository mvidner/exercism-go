package scrabble

import "strings"

// Scrabble score for the word (ASCII upper and lower case letters only, otherwise we panic)
func Score(word string) int {
	letter_values := map[byte]int{
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10,
	}

	word = strings.ToUpper(word)
	score := 0
	for i := 0; i < len(word); i++ {
		letter := word[i]
		value, ok := letter_values[letter]
		if ok {
			score += value
		} else {
			panic("scoring non-letters is not defined")
		}
	}

	return score
}
