package myutf8

func ValidRune(r rune) bool {
	if r >= 0x00 && r < 0xD800 || r > 0xDFFF && r <= 0x10FFFF { // 0xD800 -> 0xDFFF is Surrogates category
		return true
	}

	return false
}
