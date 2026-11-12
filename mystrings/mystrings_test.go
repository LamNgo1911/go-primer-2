package mystrings

import (
	"strings"
	"testing"
)

func testIndexAndCount(t *testing.T, name string, fn, standardFn func(string, string) int) {
	t.Helper()

	for _, input := range twoStringsInput {
		got := fn(input.s, input.subStr)
		expected := standardFn(input.s, input.subStr)

		if got != expected {
			t.Errorf("%s(%s, %s) = %d, want %d", name, input.s, input.subStr, got, expected)
		}

	}
}

func TestIndex(t *testing.T) {
	testIndexAndCount(t, "Index", Index, strings.Index)
}

func TestCount(t *testing.T) {
	testIndexAndCount(t, "Count", Count, strings.Count)
}

func TestHasPrefix(t *testing.T) {
	for _, input := range twoStringsInput {
		got := HasPrefix(input.s, input.subStr)
		expected := strings.HasPrefix(input.s, input.subStr)

		if got != expected {
			t.Errorf("HasPrefix(%s, %s) = %t, want %t", input.s, input.subStr, got, expected)
		}

	}
}

func TestCut(t *testing.T) {
	for _, input := range twoStringsInput {
		beforeGot, afterGot, foundGot := Cut(input.s, input.subStr)
		beforeExpected, afterExpected, foundExpected := strings.Cut(input.s, input.subStr)

		if beforeGot != beforeExpected || afterGot != afterExpected || foundGot != foundExpected {
			t.Errorf("Cut(%s, %s) = (%s, %s, %t), want (%s, %s, %t)",
				input.s, input.subStr,
				beforeGot, afterGot, foundGot,
				beforeExpected, afterExpected, foundExpected,
			)
		}

	}
}

func TestIndexRune(t *testing.T) {
	for _, input := range strAndRuneInput {
		got := IndexRune(input.s, input.r)
		expected := strings.IndexRune(input.s, input.r)

		if got != expected {
			t.Errorf("IndexRune(%s, %c) = %d, want %d", input.s, input.r, got, expected)
		}

	}
}

func TestTrimSpace(t *testing.T) {
	for _, input := range trimSpaceStrTest {
		got := TrimSpace(input.s)
		expected := input.output

		if got != expected {
			t.Errorf("TrimSpace(%s) = %s, want %s", input.s, got, expected)
		}

	}
}

func TestToLower(t *testing.T) {
	for _, input := range toLowerTest {
		got := ToLower(input.s)
		expected := input.output

		if got != expected {
			t.Errorf("ToTestToLower(%s) = %s, want %s", input.s, got, expected)
		}

	}
}

func TestRepeat(t *testing.T) {
	for _, input := range repeatTest {
		t.Run(input.name, func(t *testing.T) {
			got, gotOK := Repeat(input.s, input.count)

			if gotOK != input.expectedOk {
				t.Errorf("Repeat(%q, %d) = %t, want %t",
					input.s, input.count, gotOK, input.expectedOk)
			}

			if got != input.expectedStr {
				t.Errorf("Repeat(%q, %d) = %q, want %q",
					input.s, input.count, got, input.expectedStr)
			}
		})
	}
}
