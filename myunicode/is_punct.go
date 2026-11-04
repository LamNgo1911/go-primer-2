package myunicode

func IsPunct(r rune) bool {
	if r < '!' || r > '~' {
		return false
	}

	if IsDigit(r) || IsLetter(r) {
		return false
	}

	return true
}
