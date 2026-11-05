package myutf8

import (
	"testing"
	"unicode/utf8"
)

func TestRuneLen(t *testing.T) {
	for _, input := range runeInputs {
		got := RuneLen(input)
		expected := utf8.RuneLen(input)

		if got != expected {
			t.Errorf("RuneLen(%c) = %d, want %d", input, got, expected)
		}

	}
}
