package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var result strings.Builder
	for _, w := range Words(s) {
		result.WriteString(strings.ToUpper(FirstCharacter(w)))
	}
	return result.String()
}

func Words(s string) []string {
	isDelimiter := func(c rune) bool {
		return !unicode.IsLetter(c) && !(c == '\'')
	}
	return strings.FieldsFunc(s, isDelimiter)
}

// First character of the string, as a string
// (because s[0] returns a byte, not a rune)
func FirstCharacter(s string) string {
	for _, rune := range s {
		// exit the function in the 1st iteration
		return string(rune)
	}
	return ""
}
