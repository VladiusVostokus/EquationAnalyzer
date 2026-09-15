package lexer

import (
	"unicode"
	"golang.org/x/exp/slices"
)

type Lexer struct {
	tokens []Token
	Errors map[int]string
}

func (l *Lexer) Check(eq string) []Token {
	l.Errors = make(map[int]string)
	curLex := ""
	for i := 0; i < len(eq); i++ {
		curLex += string(eq[i])
		val, ok := GetToken[curLex]
		if ok {
			token := Token{Type: val, Value: curLex, Pos: i}
			l.tokens = append(l.tokens, token)
			curLex = ""
			continue
		} 
		if unicode.IsDigit(rune(eq[i])) {
			token := Token{Type: NUMBER, Value: curLex, Pos: i}
			val, idx := l.GetRemainingNumber(eq, i)
			i = idx
			token.Value += string(val)
			l.tokens = append(l.tokens, token)
			curLex = ""
			continue
		}
		if !slices.Contains(forbidenSymbols, curLex) {
			token := Token{Type: IDENTIFIER, Value: curLex, Pos: i}
			l.tokens = append(l.tokens, token)
			curLex = ""
			continue
		} else {
			token := Token{Type: INVALID, Value: curLex, Pos: i}
			l.tokens = append(l.tokens, token)
			l.Errors[i] = "Forbidden lexema: " + curLex
			curLex = ""
			continue
		}
	}
	return l.tokens
}

func (l *Lexer) GetRemainingNumber(eq string, start int) (string, int) {
	i := start
	val := ""
	hasDot := false
	for j := i + 1; j < len(eq); j++ {
		lexema := rune(eq[j])
		if unicode.IsDigit(lexema) {
			val += string(lexema)
				i++
			} else {

			if lexema == '.' {
				if hasDot {
					l.Errors[j] = "Unexpected dot"
				}
				hasDot = true
				val += string(lexema)
				i++
			} else {
				break
			}
		}
	}
	return val, i
}