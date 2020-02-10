package bob

import "strings"

// Bob's reply to *remark*
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	shouting := IsShouting(remark)
	question := IsQuestion(remark)
	if shouting && question {
		return "Calm down, I know what I'm doing!"
	}
	if shouting {
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	return "Whatever."
}

// There are letters, and they are all in uppercase
func IsShouting(s string) bool {
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hasLetters := strings.ContainsAny(s, upper)
	return hasLetters && s == strings.ToUpper(s)
}

func IsQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}
