package lexer

type TokenType int

const (
	PCDeclaracoes TokenType = iota
	PCAlgoritmo
	PCInteiro
	PCReal
	PCAtribuir
	PCA
	PCLer
	PCImprimir
	PCSe
	PCEntao
	PCSenao
	PCEnquanto
	PCInicio
	PCFim
	OpAritMult
	OpAritDiv
	OpAritSoma
	OpAritSub
	OpRelMenor
	OpRelMenorIgual
	OpRelMaiorIgual
	OpRelMaior
	OpRelIgual
	OpRelDif
	OpBoolE
	OpBoolOu
	Delim
	AbrePar
	FechaPar
	Var
	NumInt
	NumReal
	Cadeia
	Fim
)

func (t TokenType) String() string {
	return [...]string{
		"PCDeclaracoes",
		"PCAlgoritmo",
		"PCInteiro",
		"PCReal",
		"PCAtribuir",
		"PCA",
		"PCLer",
		"PCImprimir",
		"PCSe",
		"PCEntao",
		"PCSenao",
		"PCEnquanto",
		"PCInicio",
		"PCFim",
		"OpAritMult",
		"OpAritDiv",
		"OpAritSoma",
		"OpAritSub",
		"OpRelMenor",
		"OpRelMenorIgual",
		"OpRelMaiorIgual",
		"OpRelMaior",
		"OpRelIgual",
		"OpRelDif",
		"OpBoolE",
		"OpBoolOu",
		"Delim",
		"AbrePar",
		"FechaPar",
		"Var",
		"NumInt",
		"NumReal",
		"Cadeia",
		"Fim",
	}[t]
}
