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
	ui := &ui.UI{Lexer: lexer, Parser: parser}
	for {
		var eq string
		fmt.Println("Enter equation: ")
		fmt.Scanln(&eq)
		ui.GetInput(eq)
		ui.MakeCheck()
		fmt.Println(ui.FormOutputErrors())
		fmt.Println("=========================================")
	}
}