package myunicode

func ToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r +=  'a'
	}

	return r
}