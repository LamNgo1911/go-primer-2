package myunicode

func ToUpper(r rune) rune{
	if r >= 'a' &&  r <= 'z'{
		r -= 'a' - 'A'
	}

	return r
}