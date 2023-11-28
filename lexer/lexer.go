package lexer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Lexer struct {
	file   *os.File
	reader *bufio.Reader
}

func ReadFile(filepath string) *Lexer {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	reader := bufio.NewReader(file)
	return &Lexer{file, reader}
}

func (lexer *Lexer) readNextChar() (rune, error) {
	rune, _, err := lexer.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			fmt.Println("End of file")
		} else {
			fmt.Println("Error reading file:", err)
		}
		return rune, err
	}
	fmt.Print(string(rune))
	return rune, nil
}

func (lexer *Lexer) ReadNextToken() *Token {
	for {
		c, err := lexer.readNextChar()
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		if c == ' ' || c == '\n' {
			continue
		}

		if c == ':' {
			return &Token{Delim, ":"}
		} else if c == '*' {
			return &Token{OpAritMult, "*"}
		} else if c == '/' {
			return &Token{OpAritDiv, "/"}
		} else if c == '+' {
			return &Token{OpAritSoma, "+"}
		} else if c == '-' {
			return &Token{OpAritSub, "-"}
		} else if c == '(' {
			return &Token{AbrePar, "("}
		} else if c == ')' {
			return &Token{FechaPar, ")"}
		} else if c == '<' {
			c, _ := lexer.readNextChar()
			if c == '=' {
				return &Token{OpRelMenorIgual, "<="}
			} else if c == '>' {
				return &Token{OpRelDif, "<>"}
			} else {
				_ = lexer.reader.UnreadRune()
				return &Token{OpRelMenor, "<"}
			}
		}

	}
	return nil
}
