package lexer

type TokenType int

const (
	NUMBER TokenType = iota
	IDENTIFIER
	OPERATOR
	OP_PAREN
	CL_PAREN
	INVALID
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
	"(": OP_PAREN,
	")": CL_PAREN,
}

var forbidenSymbols = []string{"&","$","_","@","#",".","!","?","^","%"}