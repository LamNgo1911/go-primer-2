package myutf8

var runeInputs = []rune{
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

var byteInputs = []struct {
	name   string
	p      []byte
	r      rune
	output int
	ok     bool
}{
	{"Length of P shorter", []byte{0x00}, '日', 0, false},
	{"Valid byte slice length and runeLength = 1", []byte{0x00, 0x00}, 'a', 1, true},
	{"Valid byte slice length and runeLength = 3", []byte{0x00, 0x00, 0x00}, '日', 3, true},
	{"Invalid rune", []byte{0x00, 0x00, 0x00, 0x00}, 0xD800, 3, true},
}
