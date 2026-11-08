package myunicode

func IsSpace(r rune) bool {
	return (r >= 0x09 && r <= 0x0D) || r == 0x20
}
