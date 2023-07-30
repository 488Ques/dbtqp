package dbtqp

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

// Return a new instance of Scanner
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()

}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as a tag or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isTagRuneStarter(ch) {
		s.unread()
		return s.scanTag()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '-':
		return NEGATE, string(ch)
	case '(':
		return LPAR, string(ch)
	case ')':
		return RPAR, string(ch)
	}

	if ch == eof {
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhiteSpace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanTag() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	lparRead := false

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isTagRune(ch) || (ch == ')' && !lparRead) {
			s.unread()
			break
		} else {
			if ch == '(' {
				lparRead = true
			}
			buf.WriteRune(ch)
		}
	}

	if strings.ToLower(buf.String()) == "or" {
		s.scanWhiteSpace() // Skip the WS following OR
		return OR, buf.String()
	}

	return TAG, buf.String()
}
