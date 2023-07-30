package dbtqp

import (
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

// Put previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWS() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// TODO Make it accept a slice of expected tokens instead of one expected token
func WrongTokenError(wrongTok Token, lit string, tok Token) error {
	return fmt.Errorf("found %s (literal: %q), expected %s", wrongTok, lit, tok)
}

func (p *Parser) Parse() ([]TagEntity, error) {
	qr := make([]TagEntity, 0)
	negate := false

	for i := 0; ; {
		// Scan the next TAG or OR keyword
		tok, lit := p.scanIgnoreWS()

		// Out of things to scan
		if tok == EOF {
			break
		}

		if tok == RPAR {
			break
		}

		if tok == NEGATE {
			negate = true
			tok, lit = p.scanIgnoreWS()
		}

		switch tok {
		case TAG:
			t := NewTag(lit, negate)
			if i > 0 {
				qr[i-1].SetNext(t)
			}
			qr = append(qr, t)

			negate = false
			i++
		case OR:
			if i == 0 {
				return nil, WrongTokenError(tok, lit, TAG)
			}
			qr[i-1].SetRelationship("OR")
		case LPAR:
			subqr, err := p.Parse()
			if err != nil {
				return nil, err
			}

			ts := make([]*Tag, 0)
			for _, te := range subqr {
				t := te.(*Tag)
				ts = append(ts, t)
			}
			tg := NewTagGroup(ts, negate)

			qr = append(qr, tg)

			negate = false
			i++
		default:
			return nil, fmt.Errorf("found %q, expected TAG or OR", lit)
		}
	}

	return qr, nil
}

func NewParserString(query string) *Parser {
	return NewParser(strings.NewReader(query))
}
