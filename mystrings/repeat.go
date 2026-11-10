package mystrings

func Repeat(s string, count int) (string, bool)  {
	// maxInt64 := -1 >> 1
	maxInt64 := int(^uint64(0) >> 1)
	if count < 0  || len(s) >(maxInt64 / count) {
		return "", false
	}

	if count == 0 {
		return "", true
	}

	result := make([]byte, 0, len(s)*count)
	// var result []byte
	for i := 0; i < count; i++ {
		result = append(result, s...)
	}

	return string(result), true
}

// 9223372036854775807