package mystrings

func Count(s, substr string) int { // count rune

	if substr == "" {
		return len([]rune(s)) + 1
	}

	count := 0
	for i := 0; i <= len(s)-len(substr); {
		if s[i:i+len(substr)] == substr {
			count++
			i += len(substr)
		} else {
			i++
		}
	}

	return count
}
