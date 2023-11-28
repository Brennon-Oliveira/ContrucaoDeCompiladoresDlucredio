package lexer

type Token struct {
	Type   TokenType
	Lexeme string
}

func NewToken(tokenType TokenType, lexeme string) Token {
	return Token{tokenType, lexeme}
}

func (t Token) String() string {
	return "<" + t.Type.String() + ", " + t.Lexeme + ">"
}
