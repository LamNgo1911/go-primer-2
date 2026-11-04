package myutf8

func ValidRune(r rune) bool {
	if r >= 0x00 && r < 0xD800 || r > 0xDFFF && r <= 0x10FFFF {
		return true
	}

	return false
}
