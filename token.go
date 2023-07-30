package dbtqp

// Token represents a lexical token
type Token int

const (
	// Special token
	ILLEGAL Token = iota
	EOF
	WS // Whitespace

	// Literals
	TAG

	// Misc
	LPAR // (
	RPAR // )

	// Keywords
	OR     // 'or'
	NEGATE // -
)

func (tok Token) String() string {
	switch tok {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case WS:
		return "WS"
	case TAG:
		return "TAG"
	case LPAR:
		return "LPAR"
	case RPAR:
		return "RPAR"
	case OR:
		return "OR"
	case NEGATE:
		return "-"
	}

	return "Unknown"
}

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isParenthesis(ch rune) bool {
	return ch == '(' || ch == ')'
}

func isQuote(ch rune) bool {
	return ch == '"' || ch == '\''
}

// Check whether the rune is valid in a tag name
func isTagRune(ch rune) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_' || isParenthesis(ch) || isQuote(ch)
}

// Check whether the rune is valid as the first rune in a tag name
func isTagRuneStarter(ch rune) bool {
	return isLetter(ch) || isDigit(ch) || isQuote(ch)
}
