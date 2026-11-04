package mystrings

func HasPrefix(s, prefix string) bool {
	if prefix == "" {
		return true
	}

	if len(s) >= len(prefix) {
		return true
	}

	if s[0] == prefix[0] && s[:len(prefix)] == prefix {
		return true
	}

	return false
}
