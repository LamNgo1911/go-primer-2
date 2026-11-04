package goprimer

func ToUpper(r rune) rune{
	if r >= 'a' &&  r >= 'z'{
		r -= 'A'
	}

	return r
}