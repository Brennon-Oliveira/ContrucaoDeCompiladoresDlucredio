package main

import (
	LexerPackage "ContrucaoDeCompiladoresDlucredio/lexer"
	"fmt"
)

func main() {
	filePath := "./algoritmo.alguma"
	lexer := LexerPackage.ReadFile(filePath)

	for {
		token := lexer.ReadNextToken()
		if token == nil {
			break
		}
		fmt.Println(token)
	}
}
