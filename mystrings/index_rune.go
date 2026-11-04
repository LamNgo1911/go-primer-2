package mystrings

func IndexRune(s string, r rune) int {
	if r < 0 || r > 0x10FFFF || (r >= 0xD800 && r <= 0xDFFF) {
		r = -1
	}

	for i, char := range s {
		if char == r {
			return i
		}
	}

	return -1
}
