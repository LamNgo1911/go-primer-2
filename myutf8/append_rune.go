package myutf8

func AppendRune(p []byte, r rune) []byte {
	// Invalid characters default to the standard replacement character U+FFFD
	if r < 0 || r > 0x10FFFF || (r >= 0xD800 && r <= 0xDFFF) {
		r = 0xFFFD
	}

	if r <= 0x7F {
		return append(p, byte(r))
	} else if r <= 0x7FF {
		byte1 := byte(0b11000000 | r>>6)
		byte2 := byte(0x80 | r&0x3F) // 0x80 = 10000000 0x3F = 00111111

		return append(p, byte1, byte2)
	} else if r <= 0xFFFF {
		byte1 := byte(0b11100000 | r>>12)
		byte2 := byte(0x80 | (r>>6)&0x3F)
		byte3 := byte(0x80 | r&0x3F)

		return append(p, byte1, byte2, byte3)
	} else {
		byte1 := byte(0b11110000 | r>>18)
		byte2 := byte(0x80 | (r>>12)&0x3F)
		byte3 := byte(0x80 | (r>>6)&0x3F)
		byte4 := byte(0x80 | r&0x3F)

		return append(p, byte1, byte2, byte3, byte4)
	}
}
