package mystrings

func TrimSpace(s string) string {
	start := 0
	end := len(s)

	if start < end && IsASCIIWhitespace(s[start]) {
		start++
	}

	if end > start && IsASCIIWhitespace(s[end-1]) {
		end--
	}

	return s[start:end]
}

func IsASCIIWhitespace(b byte) bool {
	return b >= 0x09 && b <= 0x0D || b == 0x20
}
