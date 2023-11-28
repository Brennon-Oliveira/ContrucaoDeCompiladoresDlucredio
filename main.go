package main

import (
	"ContrucaoDeCompiladoresDlucredio/parser"
	"fmt"
	"github.com/antlr4-go/antlr"
)

func main() {
	filePath := "./algoritmo.alguma"
	println("Compilando arquivo: " + filePath)

	parser.AlgumaGrammarInit()
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewAlgumaGrammar(input)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("<%s, %s>\n", t.GetText(), lexer.SymbolicNames[t.GetTokenType()])
	}
}
