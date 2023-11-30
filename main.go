package main

import (
	"ContrucaoDeCompiladoresDlucredio/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "./sintatica.alguma"
	println("Compilando arquivo: " + filePath)

	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewAlgumaGrammarLexer(input)

	//t := lexer.NextToken()
	//for t.GetTokenType() != antlr.TokenEOF {
	//	fmt.Printf("<%s,%s>\n", t.GetText(), lexer.SymbolicNames[t.GetTokenType()])
	//	t = lexer.NextToken()
	//}

	commonTokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parserInstance := parser.NewAlgumaGrammarParser(commonTokenStream)
	parserInstance.Programa()
}
