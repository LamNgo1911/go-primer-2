package myutf8

func RuneLen(r rune) int {
	if !ValidRune(r) {
		return -1
	}

	if r < 0x80 { // 7 bits max
		return 1
	} else if r < 0x800 { // 11 bits max
		return 2
	} else if r < 0x10000 { // 17 bits max
		return 3
	} else { // r < 0x10FFFF // 21 bits max

		return 4
	}
}
