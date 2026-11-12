package myutf8

func EncodeRune(p []byte, r rune) (int, bool) {

	// Invalid characters default to the standard replacement character U+FFFD
	if r < 0 || r > 0x10FFFF || (r >= 0xD800 && r <= 0xDFFF) {
		r = 0xFFFD // rune error �
	}

	rlen := RuneLen(r)

	if len(p) < rlen {
		return 0, false
	}

	if r <= 0x7F {
		p[0] = byte(r)
	} else if r <= 0x7FF {
		p[0] = byte(0b11000000 | r>>6)
		p[1] = byte(0x80 | r&0x3F) // 0x80 = 10000000 0x3F = 00111111
	} else if r <= 0xFFFF {
		p[0] = byte(0b11100000 | r>>12)
		p[1] = byte(0x80 | (r>>6)&0x3F)
		p[2] = byte(0x80 | r&0x3F)
	} else {
		p[0] = byte(0b11110000 | r>>18)
		p[1] = byte(0x80 | (r>>12)&0x3F)
		p[2] = byte(0x80 | (r>>6)&0x3F)
		p[3] = byte(0x80 | r&0x3F)
	}

	return rlen, true
}

// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
// r = 001 111101 101000 000000
//     ^^^ ^^^^^^ ^^^^^^ ^^^^^^
//     Box0 Box1   Box2   Box3
