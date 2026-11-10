package mystrings

import "goprimer/myunicode"

func ToLower(s string) string {
	result := make([]rune, 0, len(s))
	for _, char := range s {
		result = append(result, myunicode.ToLower(char))
	}

	return string(result)
}
