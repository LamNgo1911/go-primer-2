package goprimer

func IsLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}
