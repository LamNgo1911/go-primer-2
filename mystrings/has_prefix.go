package mystrings

func HasPrefix(s, prefix string) bool {
	return len(prefix) <= len(s) && s[:len(prefix)] == prefix
}
