package twofer

// A sharing message; if the name is empty, we assume "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
