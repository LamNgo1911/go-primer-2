package myunicode

import (
	"testing"
	"unicode"
)

var inputs = []rune{
		'0',
		'7',
		'ᚢ',
		'A',
		'z',
		'ᛞ',
		'你',
	}

func TestIsDigit(t *testing.T){
	for _, input := range inputs {
		got := IsDigit(input)
		expected := unicode.IsDigit(input)

		if got != expected {
			t.Errorf("IsDigit(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestIsLetter(t *testing.T){
	isLetterInputs := []struct {
		in rune
		out bool
		}{
		{'0', false},
		{'7', false},
		{'t', true},
		{'A', true},
		{'ᚢ', false},
		{'ᛞ', false},
		{'你', false},
	}

	for _, input := range isLetterInputs {
		got := IsLetter(input.in)
		expected := input.out

		if got != expected {
			t.Errorf("IsLetter(%c) = %t, want %t", input.in, got, expected)
		}
	}
}

func TestIsSpace(t *testing.T){
	for _, input := range inputs {
		got := IsSpace(input)
		expected := unicode.IsSpace(input)

		if got != expected {
			t.Errorf("IsSpace(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestIsUpper(t *testing.T){
	for _, input := range inputs {
		got := IsUpper(input)
		expected := unicode.IsUpper(input)

		if got != expected {
			t.Errorf("IsUpper(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestIsLower(t *testing.T){
	for _, input := range inputs {
		got := IsLower(input)
		expected := unicode.IsLower(input)

		if got != expected {
			t.Errorf("IsLower(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestIsPunct(t *testing.T){
	for _, input := range inputs {
		got := IsPunct(input)
		expected := unicode.IsPunct(input)

		if got != expected {
			t.Errorf("IsPunct(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestIsControl(t *testing.T){
	for _, input := range inputs {
		got := IsControl(input)
		expected := unicode.IsControl(input)

		if got != expected {
			t.Errorf("IsControl(%c) = %t, want %t", input, got, expected)
		}
	}
}

func TestToUpper(t *testing.T){
	for _, input := range inputs {
		got := ToUpper(input)
		expected := unicode.ToUpper(input)

		if got != expected {
			t.Errorf("ToUpper(%c) = %c, want %c", input, got, expected)
		}
	}
}

func TestToLower(t *testing.T){
	for _, input := range inputs {
		got := ToLower(input)
		expected := unicode.ToLower(input)

		if got != expected {
			t.Errorf("ToLower(%c) = %c, want %c", input, got, expected)
		}
	}
}