package myutf8

func AppendRune(p []byte, r rune) []byte {
	// Invalid characters default to the standard replacement character U+FFFD
	if r < 0x00 || r > 0x10FFFF || (r >= 0xD800 && r <= 0xDFFF) {
		r = 0xFFFD
	}

	rlen := RuneLen(r)
	start := len(p)

	p = append(p, make([]byte, rlen)...)

	EncodeRune(p[start:], r)

	return p
}
