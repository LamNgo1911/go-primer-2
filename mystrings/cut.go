package mystrings

func Cut(s, sep string) (before, after string, found bool) {
	if sep == "" {
		return "", s, true
	}

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return s[:i], s[i+len(sep):], true
		}
	}

	return s, "", false
}
