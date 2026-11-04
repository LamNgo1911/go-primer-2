package myutf8

func RuneLen(r rune) int {
	if r < 0x80 {
		return 1
	} else if r&0b11100000 == 0b11000000 {
		return 2
	} else if r&0b11110000 == 0b11100000 {
		return 3
	} else if r&0b11111000 == 0b11110000 {
		return 4
	}

	return 0
}
