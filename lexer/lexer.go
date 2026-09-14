package lexer

import "fmt"

type Lexer struct {
	tokens []Token
}

func (l *Lexer) Check(eq string) []Token {
	for i, v := range eq {
		curLex := string(v)
		fmt.Println(i)
		val, ok := GetToken[curLex]
		if ok {
			token := Token{Type: val, Value: curLex, Pos: i}
			l.tokens = append(l.tokens, token)
		}
	}
	return l.tokens
}