package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = iota
	EOF

	IDENTIFIER
	INTEGER

	ASSIGN
	PLUS

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET
)

func (tt TokenType) String() string {
	switch tt {
	case ILLEGAL:
		return "<illegal>"
	case EOF:
		return "<eof>"
	case IDENTIFIER:
		return "<identifier>"
	case INTEGER:
		return "<integer>"
	case ASSIGN:
		return "<assign>"
	case PLUS:
		return "<plus>"
	case COMMA:
		return "<comma>"
	case SEMICOLON:
		return "<semicolon>"
	case LPAREN:
		return "<lparen>"
	case RPAREN:
		return "<rparen>"
	case LBRACE:
		return "<lbrace>"
	case RBRACE:
		return "<rbrace>"
	case FUNCTION:
		return "<function>"
	case LET:
		return "<let>"
	}
	return "wut?"
}
