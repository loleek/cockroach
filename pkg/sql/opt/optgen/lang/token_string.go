// Code generated by "stringer -type=Token scanner.go"; DO NOT EDIT.

package lang

import "strconv"

const _Token_name = "ILLEGALERROREOFIDENTSTRINGNUMBERWHITESPACECOMMENTLPARENRPARENLBRACKETRBRACKETLBRACERBRACEDOLLARCOLONASTERISKEQUALSARROWAMPERSANDCOMMACARETELLIPSESPIPE"

var _Token_index = [...]uint8{0, 7, 12, 15, 20, 26, 32, 42, 49, 55, 61, 69, 77, 83, 89, 95, 100, 108, 114, 119, 128, 133, 138, 146, 150}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}