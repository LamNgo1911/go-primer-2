package mystrings

var twoStringsInput = []struct {
	s      string
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

var strAndRuneInput = []struct {
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

var trimSpaceStrTest = []struct {
	s      string
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

var repeatTest = []struct {
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
}
