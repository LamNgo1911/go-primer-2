package mystrings

import "strings"
import "testing"

var twoStringsInput = []struct{
	s string
	subStr string
}{
	{"mystrings", ""},
	{"", ""},
	{"", "mystrings"},
	{"mystrings", "strings"},
	{"mystri你ngs", "你"},
	{"mystri你ngs", "ngs"},
	{"mystrings", "mystrings"},
	{"mystrings", "M"},
	{"banana", "an"},
	{"hello world", " "},
	{"😀😁😂😃", "😂"},
	{"mystrings", "stringss"},
	{"abc\n\tdef", "\n"},
}

var strAndRuneInput = []struct{
	s string
	r rune
}{
	{" my strings", 't'},
	{"", ' '},
	{"", 'A'},
	{"mystrings", 0xD800},
	{"    mystri你ngs", '你'},
	{"mys tri你ng  s", 'n'},
	{"myst   rings", '😂'},
	{"😀😁😂😃", '😂'},
	{"mystring  s∇ ", '∇'},
	{"abc\n\t def", '∇'},
}

var trimSpaceStrTest = []struct{
	s string
	output string
}{
	{" my strings", "my strings"},
	{"", ""},
	{"mystrings           ", "mystrings"},
	{"    mystri你  ngs   ", "mystri你  ngs"},
	{"😀😁😂    😃", "😀😁😂    😃"},
	{"mystring  s∇ ", "mystring  s∇"},
	{"abc\n\t def\n\t\n\t   ", "abc\n\t def"},
}

var toLowerTest = []struct {
	s      string
	output string
}{
	{" MY STRINGS", " my strings"},
	{"", ""},
	{"MYSTRINGS           ", "mystrings           "},
	{"    MYSTRI你  NGS   ", "    mystri你  ngs   "},
	{"😀😁😂    😃", "😀😁😂    😃"},
	{"MYSTRING  S∇ ", "mystring  s∇ "},
	{"ABC\n\t DEF\n\t\n\t   ", "abc\n\t def\n\t\n\t   "},
}

var repeatTest = []struct{
	name        string
	s           string
	count       int
	expectedStr string
	expectedOk  bool
}{
	{
	name:        "Standard repetition",
	s:           "abc",
	count:       3,
	expectedStr: "abcabcabc",
	expectedOk:  true,
},
{
	name:        "Single repetition",
	s:           "Go",
	count:       1,
	expectedStr: "Go",
	expectedOk:  true,
},
{
	name:        "Repetition with spaces and newlines",
	s:           "a\n ",
	count:       2,
	expectedStr: "a\n a\n ",
	expectedOk:  true,
},
{
	name:        "Repetition with Unicode/Runes",
	s:           "😀你",
	count:       2,
	expectedStr: "😀你😀你",
	expectedOk:  true,
},
{
	name:        "Empty string with positive count",
	s:           "",
	count:       5,
	expectedStr: "",
	expectedOk:  true,
},
{
	name:        "Zero count",
	s:           "abc",
	count:       0,
	expectedStr: "",
	expectedOk:  true,
},
{
	name:        "Negative count",
	s:           "abc",
	count:       -1,
	expectedStr: "",
	expectedOk:  false,
},
{
	name:        "Potential overflow count",
	s:           "large_string",
	count:       1e9,
	expectedStr: "",
	expectedOk:  false,
},
}

func testIndexAndCount(t *testing.T, name string, fn, standardFn func(string, string) int)  {
	t.Helper()

	for _, input := range twoStringsInput {
		got := fn(input.s, input.subStr)
		expected := standardFn(input.s, input.subStr)

		if got != expected{
			t.Errorf("%s(%s, %s) = %d, want %d", name, input.s, input.subStr, got, expected)
		}
		
	}
}

func TestIndex(t *testing.T){
	testIndexAndCount(t, "Index", Index, strings.Index)
}

func TestCount(t *testing.T){
	testIndexAndCount(t, "Count", Count, strings.Count)
}

func TestHasPrefix(t *testing.T){
	for _, input := range twoStringsInput {
		got := HasPrefix(input.s, input.subStr)
		expected := strings.HasPrefix(input.s, input.subStr)

		if got != expected {
			t.Errorf("HasPrefix(%s, %s) = %t, want %t", input.s, input.subStr,got, expected)
		}

	}
}

func TestCut(t *testing.T){
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

func TestIndexRune(t *testing.T){
	for _, input := range strAndRuneInput {
		got := IndexRune(input.s, input.r)
		expected := strings.IndexRune(input.s, input.r)

		if got != expected {
			t.Errorf("IndexRune(%s, %c) = %d, want %d", input.s, input.r,got, expected)
		}

	}
}

func TestTrimSpace(t *testing.T){
	for _, input := range trimSpaceStrTest {
		got := TrimSpace(input.s)
		expected := input.output

		if got != expected {
			t.Errorf("TrimSpace(%s) = %s, want %s", input.s ,got, expected)
		}

	}
}

func TestToLower(t *testing.T){
	for _, input := range toLowerTest {
		got := ToLower(input.s)
		expected := input.output

		if got != expected {
			t.Errorf("ToTestToLower(%s) = %s, want %s", input.s ,got, expected)
		}

	}
}

func TestRepeat(t *testing.T) {
	for _, input := range repeatTest {
		t.Run(input.name, func(t *testing.T) {
			got, gotOK := Repeat(input.s, input.count)
			expectedStr := input.expectedStr
			expectedOk := input.expectedOk

			if gotOK != expectedOk {
				t.Errorf("Repeat(%s, %d) = %t, want %t",
					input.s, input.count, gotOK, expectedOk)
			}

			if got != expectedStr {
				t.Errorf("Repeat(%s, %d) = %s, want %s",
					input.s, input.count, got, expectedStr)
			}
		})
	}
}




