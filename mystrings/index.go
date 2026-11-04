package mystrings

func Index(s, substr string) int { // count byte
	if substr == "" {
		return 0
	}

	if s == "" {
		return -1
	}

	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}

	return -1

}

// 0-9
// 7-11  i <= s - substr
