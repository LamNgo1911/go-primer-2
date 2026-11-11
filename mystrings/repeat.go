package mystrings

func Repeat(s string, count int) (string, bool) {
	if count < 0 {
		return "", false
	}

	if count == 0 || len(s) == 0 {
		return "", true
	}

	// maxInt64 := -1 >> 1
	maxInt := int(^uint(0) >> 1)
	if count > (maxInt / len(s)) {
		return "", false
	}

	result := make([]byte, 0, len(s)*count)
	// var result []byte
	for i := 0; i < count; i++ {
		result = append(result, s...)
	}

	return string(result), true
}

// 9223372036854775807
