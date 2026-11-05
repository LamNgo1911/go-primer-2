package myutf8

func RuneLen(r rune) int {
	if !ValidRune(r) {
		return -1
	}

	if r < 0x80 {
		return 1
	} else if r < 0x800 {
		return 2
	} else if r < 0x10000 {
		return 3
	} else { // r < 0x10FFFF
		return 4
	}
}
