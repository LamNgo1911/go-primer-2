package myutf8

import (
	"testing"
	"unicode/utf8"
)

func TestValidRune(t *testing.T) {
	for _, input := range inputs {
		got := ValidRune(input)
		expected := utf8.ValidRune(input)

		if got != expected {
			t.Errorf("ValidRune(%c) = %t, want %t", input, got, expected)
		}

	}
}
