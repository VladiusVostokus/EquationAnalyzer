package lexer

import (
	"unicode"
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
				val, idx := l.GetRemainingNumber(eq, i)
				i = idx
				token.Value += string(val)
				l.tokens = append(l.tokens, token)
				curLex = ""
			}
		}
	}
	return l.tokens
}

func (l *Lexer) GetRemainingNumber(eq string, start int) (string, int) {
	i := start
	val := ""
	for j := i + 1; j < len(eq); j++ {
		lexema := rune(eq[j])
		if unicode.IsDigit(lexema) {
			val += string(lexema)
				i++
			} else {

			if lexema == '.' {
				val += string(lexema)
				i++
			} else {
				break
			}
		}
	}
	return val, i
}