package myunicode

func IsControl(r rune) bool {
	return r < 0x20 || r == 0x7f // Del
}
