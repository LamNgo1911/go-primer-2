package myutf8

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestAppendRune(t *testing.T) {
	var p []byte
	for _, input := range runeInputs {
		got := AppendRune(p, input)
		expected := utf8.AppendRune(p, input)

		if !bytes.Equal(got, expected) {
			t.Errorf("AppendRune(% X, U+%04X) = % X, want % X",
				p, input, got, expected)
		}

	}
}
