// Code generated from ./parser/AlgumaGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AlgumaGrammar struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AlgumaGrammarLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algumagrammarLexerInit() {
	staticData := &AlgumaGrammarLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "':'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "PalavraChave", "NumInt", "NumReal", "Variavel", "Cadeia", "OP_REL",
		"OP_ARIT", "DELIM", "ABREPAR", "FECHAPAR", "WS", "Letra", "Digito",
	}
	staticData.RuleNames = []string{
		"PalavraChave", "NumInt", "NumReal", "Variavel", "Cadeia", "ESC_SEQ",
		"OP_REL", "OP_ARIT", "DELIM", "ABREPAR", "FECHAPAR", "WS", "Letra",
		"Digito",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 179, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 3, 0, 108, 8, 0, 1, 1, 3, 1, 111, 8, 1, 1, 1, 4, 1, 114,
		8, 1, 11, 1, 12, 1, 115, 1, 2, 3, 2, 119, 8, 2, 1, 2, 4, 2, 122, 8, 2,
		11, 2, 12, 2, 123, 1, 2, 1, 2, 4, 2, 128, 8, 2, 11, 2, 12, 2, 129, 1, 3,
		1, 3, 1, 3, 5, 3, 135, 8, 3, 10, 3, 12, 3, 138, 9, 3, 1, 4, 1, 4, 1, 4,
		5, 4, 143, 8, 4, 10, 4, 12, 4, 146, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 162, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 0, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 1,
		0, 6, 2, 0, 43, 43, 45, 45, 1, 0, 92, 92, 2, 0, 60, 60, 62, 62, 3, 0, 42,
		43, 45, 45, 47, 47, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122,
		204, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 110, 1, 0, 0, 0,
		5, 118, 1, 0, 0, 0, 7, 131, 1, 0, 0, 0, 9, 139, 1, 0, 0, 0, 11, 149, 1,
		0, 0, 0, 13, 161, 1, 0, 0, 0, 15, 163, 1, 0, 0, 0, 17, 165, 1, 0, 0, 0,
		19, 167, 1, 0, 0, 0, 21, 169, 1, 0, 0, 0, 23, 171, 1, 0, 0, 0, 25, 175,
		1, 0, 0, 0, 27, 177, 1, 0, 0, 0, 29, 30, 5, 68, 0, 0, 30, 31, 5, 69, 0,
		0, 31, 32, 5, 67, 0, 0, 32, 33, 5, 76, 0, 0, 33, 34, 5, 65, 0, 0, 34, 35,
		5, 82, 0, 0, 35, 36, 5, 65, 0, 0, 36, 37, 5, 67, 0, 0, 37, 38, 5, 79, 0,
		0, 38, 39, 5, 69, 0, 0, 39, 108, 5, 83, 0, 0, 40, 41, 5, 65, 0, 0, 41,
		42, 5, 76, 0, 0, 42, 43, 5, 71, 0, 0, 43, 44, 5, 79, 0, 0, 44, 45, 5, 82,
		0, 0, 45, 46, 5, 73, 0, 0, 46, 47, 5, 84, 0, 0, 47, 48, 5, 77, 0, 0, 48,
		108, 5, 79, 0, 0, 49, 50, 5, 73, 0, 0, 50, 51, 5, 78, 0, 0, 51, 52, 5,
		84, 0, 0, 52, 53, 5, 69, 0, 0, 53, 54, 5, 73, 0, 0, 54, 55, 5, 82, 0, 0,
		55, 108, 5, 79, 0, 0, 56, 57, 5, 82, 0, 0, 57, 58, 5, 69, 0, 0, 58, 59,
		5, 65, 0, 0, 59, 108, 5, 76, 0, 0, 60, 61, 5, 65, 0, 0, 61, 62, 5, 84,
		0, 0, 62, 63, 5, 82, 0, 0, 63, 64, 5, 73, 0, 0, 64, 65, 5, 66, 0, 0, 65,
		66, 5, 85, 0, 0, 66, 67, 5, 73, 0, 0, 67, 108, 5, 82, 0, 0, 68, 108, 5,
		65, 0, 0, 69, 70, 5, 76, 0, 0, 70, 71, 5, 69, 0, 0, 71, 108, 5, 82, 0,
		0, 72, 73, 5, 73, 0, 0, 73, 74, 5, 77, 0, 0, 74, 75, 5, 80, 0, 0, 75, 76,
		5, 82, 0, 0, 76, 77, 5, 73, 0, 0, 77, 78, 5, 77, 0, 0, 78, 79, 5, 73, 0,
		0, 79, 108, 5, 82, 0, 0, 80, 81, 5, 83, 0, 0, 81, 108, 5, 69, 0, 0, 82,
		83, 5, 69, 0, 0, 83, 84, 5, 78, 0, 0, 84, 85, 5, 84, 0, 0, 85, 86, 5, 65,
		0, 0, 86, 108, 5, 79, 0, 0, 87, 88, 5, 69, 0, 0, 88, 89, 5, 78, 0, 0, 89,
		90, 5, 81, 0, 0, 90, 91, 5, 85, 0, 0, 91, 92, 5, 65, 0, 0, 92, 93, 5, 78,
		0, 0, 93, 94, 5, 84, 0, 0, 94, 108, 5, 79, 0, 0, 95, 96, 5, 73, 0, 0, 96,
		97, 5, 78, 0, 0, 97, 98, 5, 73, 0, 0, 98, 99, 5, 67, 0, 0, 99, 100, 5,
		73, 0, 0, 100, 108, 5, 79, 0, 0, 101, 102, 5, 70, 0, 0, 102, 103, 5, 73,
		0, 0, 103, 108, 5, 77, 0, 0, 104, 108, 5, 69, 0, 0, 105, 106, 5, 79, 0,
		0, 106, 108, 5, 85, 0, 0, 107, 29, 1, 0, 0, 0, 107, 40, 1, 0, 0, 0, 107,
		49, 1, 0, 0, 0, 107, 56, 1, 0, 0, 0, 107, 60, 1, 0, 0, 0, 107, 68, 1, 0,
		0, 0, 107, 69, 1, 0, 0, 0, 107, 72, 1, 0, 0, 0, 107, 80, 1, 0, 0, 0, 107,
		82, 1, 0, 0, 0, 107, 87, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0, 107, 101, 1,
		0, 0, 0, 107, 104, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 2, 1, 0, 0, 0,
		109, 111, 7, 0, 0, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		113, 1, 0, 0, 0, 112, 114, 3, 27, 13, 0, 113, 112, 1, 0, 0, 0, 114, 115,
		1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 4, 1, 0, 0,
		0, 117, 119, 7, 0, 0, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		121, 1, 0, 0, 0, 120, 122, 3, 27, 13, 0, 121, 120, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 127, 5, 46, 0, 0, 126, 128, 3, 27, 13, 0, 127, 126, 1, 0, 0,
		0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		6, 1, 0, 0, 0, 131, 136, 3, 25, 12, 0, 132, 135, 3, 25, 12, 0, 133, 135,
		3, 27, 13, 0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1,
		0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 8, 1, 0, 0, 0,
		138, 136, 1, 0, 0, 0, 139, 144, 5, 39, 0, 0, 140, 143, 3, 11, 5, 0, 141,
		143, 8, 1, 0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 146,
		1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 1, 0,
		0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 5, 39, 0, 0, 148, 10, 1, 0, 0, 0,
		149, 150, 5, 92, 0, 0, 150, 151, 5, 39, 0, 0, 151, 12, 1, 0, 0, 0, 152,
		162, 7, 2, 0, 0, 153, 154, 5, 62, 0, 0, 154, 162, 5, 61, 0, 0, 155, 156,
		5, 60, 0, 0, 156, 162, 5, 61, 0, 0, 157, 158, 5, 61, 0, 0, 158, 162, 5,
		61, 0, 0, 159, 160, 5, 60, 0, 0, 160, 162, 5, 62, 0, 0, 161, 152, 1, 0,
		0, 0, 161, 153, 1, 0, 0, 0, 161, 155, 1, 0, 0, 0, 161, 157, 1, 0, 0, 0,
		161, 159, 1, 0, 0, 0, 162, 14, 1, 0, 0, 0, 163, 164, 7, 3, 0, 0, 164, 16,
		1, 0, 0, 0, 165, 166, 5, 58, 0, 0, 166, 18, 1, 0, 0, 0, 167, 168, 5, 40,
		0, 0, 168, 20, 1, 0, 0, 0, 169, 170, 5, 41, 0, 0, 170, 22, 1, 0, 0, 0,
		171, 172, 7, 4, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 6, 11, 0, 0, 174,
		24, 1, 0, 0, 0, 175, 176, 7, 5, 0, 0, 176, 26, 1, 0, 0, 0, 177, 178, 2,
		48, 57, 0, 178, 28, 1, 0, 0, 0, 12, 0, 107, 110, 115, 118, 123, 129, 134,
		136, 142, 144, 161, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AlgumaGrammarInit initializes any static state used to implement AlgumaGrammar. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAlgumaGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlgumaGrammarInit() {
	staticData := &AlgumaGrammarLexerStaticData
	staticData.once.Do(algumagrammarLexerInit)
}

// NewAlgumaGrammar produces a new lexer instance for the optional input antlr.CharStream.
func NewAlgumaGrammar(input antlr.CharStream) *AlgumaGrammar {
	AlgumaGrammarInit()
	l := new(AlgumaGrammar)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AlgumaGrammarLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AlgumaGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlgumaGrammar tokens.
const (
	AlgumaGrammarPalavraChave = 1
	AlgumaGrammarNumInt       = 2
	AlgumaGrammarNumReal      = 3
	AlgumaGrammarVariavel     = 4
	AlgumaGrammarCadeia       = 5
	AlgumaGrammarOP_REL       = 6
	AlgumaGrammarOP_ARIT      = 7
	AlgumaGrammarDELIM        = 8
	AlgumaGrammarABREPAR      = 9
	AlgumaGrammarFECHAPAR     = 10
	AlgumaGrammarWS           = 11
	AlgumaGrammarLetra        = 12
	AlgumaGrammarDigito       = 13
)
