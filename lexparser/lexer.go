package lexparser

import (
	"fmt"
	"regexp"
)

// TokenType is an enum defining the type of token being lexed.
type TokenType int

const (
	num TokenType = iota
	variable
	plus
	minus
	times
	div
	expt
	lparen
	rparen
)

func (t TokenType) String() string {
	switch t {
	case num:
		return "Num"
	case variable:
		return "Var"
	case plus:
		return "Plus"
	case minus:
		return "Minus"
	case times:
		return "Times"
	case div:
		return "Div"
	case expt:
		return "Expt"
	case lparen:
		return "LParen"
	case rparen:
		return "RParen"
	}
	panic("lexer: invalid tokentype passed to String()")
}

// Token contains a TokenType and the Data associated with that token. It can
// then be parsed by the parser.
type Token struct {
	Type TokenType
	Data string
}

func (t Token) String() string {
	result := fmt.Sprintf("%v", t.Type)
	if t.Type == num || t.Type == variable {
		result += fmt.Sprintf(": %v", t.Data)
	}
	return result
}

// LexString takes text and returns a slice of Tokens. It applies a series of
// regex expressions to find the correct rule to apply. Currently, this is
// hardcoded, but it could be put into a file later.
func LexString(text string) []Token {
	regexMap := map[TokenType]*regexp.Regexp{
		num:      regexp.MustCompile(`\A[-+]?[0-9]*\.?[0-9]+`),
		variable: regexp.MustCompile(`\A\pL+`),
		plus:     regexp.MustCompile(`\A\+`),
		minus:    regexp.MustCompile(`\A-`),
		times:    regexp.MustCompile(`\A\*`),
		div:      regexp.MustCompile(`\A/`),
		expt:     regexp.MustCompile(`\A\^`),
		lparen:   regexp.MustCompile(`\A\(`),
		rparen:   regexp.MustCompile(`\A\)`),
	}
	bytes := []byte(text)

	var result []Token
	var found []byte

	currentPos := 0
	for currentPos < len(text) {
		found = nil
		for typ, re := range regexMap {
			found = re.Find(bytes[currentPos:])
			if found != nil {
				result = append(result, Token{Type: typ, Data: string(found)})
				currentPos += len(found)
				break
			}
		}
		if found == nil {
			// We just ignore the byte.
			currentPos++
		}
	}
	return result
}
