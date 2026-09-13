package lexer

type TokenType int

const (
	NUMBER TokenType = iota
	IDENTIFIER
	OPERATOR
	OP_PAREN
	CL_PAREN
	DOT
)

type Token struct {
	Type TokenType
	Value string
	Pos int
}

var GetToken map[string]TokenType = map[string]TokenType {
	"+": OPERATOR,
	"-": OPERATOR,
	"*": OPERATOR,
	"/": OPERATOR,
	".": DOT,
	"(": OP_PAREN,
	")": CL_PAREN,
}