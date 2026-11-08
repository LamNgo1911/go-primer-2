package mystrings

func Repeat(s string, count int) (string, bool) {
	if count < 0 {
		return "", false
	}

	if count == 0 {
		return "", true
	}

	result := make([]byte, 0, len(s)*count)
	for i := 0; i < count; i++ {
		result = append(result, s...)
	}

	return string(result), true
}
