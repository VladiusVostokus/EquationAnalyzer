package main

import (
	l "lex_an/lexer"
	p "lex_an/parser"
	ui "lex_an/ui"
	"fmt"
)

func main() {
	lexer := &l.Lexer{}
	parser := &p.EquationParser{}
	equation := "/-(++a+b+*(&c/d)*3.4-(15asb+10)+"
	ui := &ui.UI{Lexer: lexer, Parser: parser}
	ui.GetInput(equation)
	ui.MakeCheck()
	fmt.Println(ui.FormOutputErrors())
}