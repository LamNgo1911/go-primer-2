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
	'+',
	' ',
	'∇', // Delete
	'=',
	0xD800,
	0x80,
	0x08,
	0x21,
}

// helper func name can not start with Test..., Go will recognize it as Test func
func testRuneFunc(t *testing.T, name string, fn, standardFunc func(rune) bool) {
	t.Helper() // Tell testing framework the curreent func is not the actual test func

	for _, input := range inputs {
		got := fn(input)
		expected := standardFunc(input)

		if got != expected {
			t.Errorf("%s(%c) = %t, want %t", name, input, got, expected)
		}
	}
}

func TestIsDigit(t *testing.T) {
	testRuneFunc(t, "IsDigit", IsDigit, unicode.IsDigit)
}
func TestIsSpace(t *testing.T) {
	testRuneFunc(t, "IsSpace", IsSpace, unicode.IsSpace)
}

func TestIsUpper(t *testing.T) {
	testRuneFunc(t, "IsUpper", IsUpper, unicode.IsUpper)
}

func TestIsLower(t *testing.T) {
	testRuneFunc(t, "IsLower", IsLower, unicode.IsLower)
}

func testASCIIRuneFunc(t *testing.T, tests []struct {
	in  rune
	out bool
}, name string, fn func(rune) bool) {
	t.Helper() // Tell testing framework the curreent func is not the actual test func

	for _, test := range tests {
		got := fn(test.in)
		expected := test.out

		if got != expected {
			t.Errorf("%s(%c) = %t, want %t", name, test.in, got, expected)
		}
	}
}

func TestIsLetter(t *testing.T) {
	isLetterInputs := []struct {
		in  rune
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
	testASCIIRuneFunc(t, isLetterInputs, "IsLetter", IsLetter)
}

func TestIsPunct(t *testing.T) {
	isPunctInputs := []struct {
		in  rune
		out bool
	}{
		{'.', true},
		{'+', true},
		{'a', false},
		{' ', false},
		{0x00A1, false},
		{'$', true},
		{'你', false},
	}

	testASCIIRuneFunc(t, isPunctInputs, "IsPunct", IsPunct)
}

func TestIsControl(t *testing.T) {
	isControlInputs := []struct {
		in  rune
		out bool
	}{
		{'.', false},
		{' ', false},
		{0x19, true},
		{0x00, true},
		{0x01, true},
		{'$', false},
		{'你', false},
	}

	testASCIIRuneFunc(t, isControlInputs, "IsControl", IsControl)
}

func testRuneTransform(t *testing.T, name string, fn, standardFunc func(rune) rune) {
	t.Helper() // Tell testing framework the curreent func is not the actual test func

	for _, input := range inputs {
		got := fn(input)
		expected := standardFunc(input)

		if got != expected {
			t.Errorf("%s(%c) = %c, want %c", name, input, got, expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	testRuneTransform(t, "ToUpper", ToUpper, unicode.ToUpper)
}

func TestToLower(t *testing.T) {
	testRuneTransform(t, "ToLower", ToLower, unicode.ToLower)
}
