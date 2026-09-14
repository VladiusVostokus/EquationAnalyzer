package lexer

import (
	"unicode"
	"fmt"
)

type Lexer struct {
	tokens []Token
}

func (l *Lexer) Check(eq string) []Token {
	curLex := ""
	for i := 0; i < len(eq); i++ {
		curLex += string(eq[i])
		val, ok := GetToken[curLex]
		if ok {
			token := Token{Type: val, Value: curLex, Pos: i}
			l.tokens = append(l.tokens, token)
			curLex = ""
		} else {
			if unicode.IsDigit(rune(eq[i])) {
				token := Token{Type: NUMBER, Value: curLex, Pos: i}
				for j := i + 1; j < len(eq); j++ {
					lexema := rune(eq[j])
					if unicode.IsDigit(lexema) {
						token.Value += string(lexema)
						i++
					} else {
						break
					}
				}
				l.tokens = append(l.tokens, token)
			}
		}
	}
	fmt.Println(len(l.tokens))
	return l.tokens
}