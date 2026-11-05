package myunicode

func IsPunct(r rune) bool {
	if r < '!' || r > '~' {
		return false
	}

	if IsDigit(r) || IsLetter(r) {
		return false
	}

	return true
}

// For standard package, Punctuation includes Standard Punctuation, Brackets & Quotes,
// Dashes & Hyphens

// Math Operators, Currency Signs, Common Symbols are recognized as Symbols
