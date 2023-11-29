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

type AlgumaGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AlgumaGrammarLexerLexerStaticData struct {
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

func algumagrammarlexerLexerInit() {
	staticData := &AlgumaGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'DECLARACOES'", "'ALGORITMO'", "'('", "')'", "'ATRIBUIR'",
		"'A'", "'LER'", "'IMPRIMIR'", "'SE'", "'ENTAO'", "'SENAO'", "'ENQUANTO'",
		"'INICIO'", "'FIM'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TIPO_VAR",
		"NUMINT", "NUMREAL", "VARIAVEL", "CADEIA", "OP_ARIT1", "OP_ARIT2", "OP_REL",
		"OP_BOOL", "COMENTARIO", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "TIPO_VAR", "NUMINT",
		"NUMREAL", "VARIAVEL", "CADEIA", "OP_ARIT1", "OP_ARIT2", "OP_REL", "OP_BOOL",
		"ESC_SEQ", "COMENTARIO", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 229, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		154, 8, 15, 1, 16, 4, 16, 157, 8, 16, 11, 16, 12, 16, 158, 1, 17, 4, 17,
		162, 8, 17, 11, 17, 12, 17, 163, 1, 17, 1, 17, 4, 17, 168, 8, 17, 11, 17,
		12, 17, 169, 3, 17, 172, 8, 17, 1, 18, 1, 18, 5, 18, 176, 8, 18, 10, 18,
		12, 18, 179, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 184, 8, 19, 10, 19, 12,
		19, 187, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 204, 8, 22, 1,
		23, 1, 23, 1, 23, 3, 23, 209, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		5, 25, 216, 8, 25, 10, 25, 12, 25, 219, 9, 25, 1, 25, 3, 25, 222, 8, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 0, 51, 25, 53, 26, 1, 0, 7, 2, 0, 65, 90, 97, 122, 3, 0,
		48, 57, 65, 90, 97, 122, 2, 0, 39, 39, 92, 92, 2, 0, 43, 43, 45, 45, 2,
		0, 42, 42, 47, 47, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 243,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0,
		3, 57, 1, 0, 0, 0, 5, 69, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 81, 1, 0, 0,
		0, 11, 83, 1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 94, 1, 0, 0, 0, 17, 98,
		1, 0, 0, 0, 19, 107, 1, 0, 0, 0, 21, 110, 1, 0, 0, 0, 23, 116, 1, 0, 0,
		0, 25, 122, 1, 0, 0, 0, 27, 131, 1, 0, 0, 0, 29, 138, 1, 0, 0, 0, 31, 153,
		1, 0, 0, 0, 33, 156, 1, 0, 0, 0, 35, 161, 1, 0, 0, 0, 37, 173, 1, 0, 0,
		0, 39, 180, 1, 0, 0, 0, 41, 190, 1, 0, 0, 0, 43, 192, 1, 0, 0, 0, 45, 203,
		1, 0, 0, 0, 47, 208, 1, 0, 0, 0, 49, 210, 1, 0, 0, 0, 51, 213, 1, 0, 0,
		0, 53, 226, 1, 0, 0, 0, 55, 56, 5, 58, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58,
		5, 68, 0, 0, 58, 59, 5, 69, 0, 0, 59, 60, 5, 67, 0, 0, 60, 61, 5, 76, 0,
		0, 61, 62, 5, 65, 0, 0, 62, 63, 5, 82, 0, 0, 63, 64, 5, 65, 0, 0, 64, 65,
		5, 67, 0, 0, 65, 66, 5, 79, 0, 0, 66, 67, 5, 69, 0, 0, 67, 68, 5, 83, 0,
		0, 68, 4, 1, 0, 0, 0, 69, 70, 5, 65, 0, 0, 70, 71, 5, 76, 0, 0, 71, 72,
		5, 71, 0, 0, 72, 73, 5, 79, 0, 0, 73, 74, 5, 82, 0, 0, 74, 75, 5, 73, 0,
		0, 75, 76, 5, 84, 0, 0, 76, 77, 5, 77, 0, 0, 77, 78, 5, 79, 0, 0, 78, 6,
		1, 0, 0, 0, 79, 80, 5, 40, 0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 41, 0, 0,
		82, 10, 1, 0, 0, 0, 83, 84, 5, 65, 0, 0, 84, 85, 5, 84, 0, 0, 85, 86, 5,
		82, 0, 0, 86, 87, 5, 73, 0, 0, 87, 88, 5, 66, 0, 0, 88, 89, 5, 85, 0, 0,
		89, 90, 5, 73, 0, 0, 90, 91, 5, 82, 0, 0, 91, 12, 1, 0, 0, 0, 92, 93, 5,
		65, 0, 0, 93, 14, 1, 0, 0, 0, 94, 95, 5, 76, 0, 0, 95, 96, 5, 69, 0, 0,
		96, 97, 5, 82, 0, 0, 97, 16, 1, 0, 0, 0, 98, 99, 5, 73, 0, 0, 99, 100,
		5, 77, 0, 0, 100, 101, 5, 80, 0, 0, 101, 102, 5, 82, 0, 0, 102, 103, 5,
		73, 0, 0, 103, 104, 5, 77, 0, 0, 104, 105, 5, 73, 0, 0, 105, 106, 5, 82,
		0, 0, 106, 18, 1, 0, 0, 0, 107, 108, 5, 83, 0, 0, 108, 109, 5, 69, 0, 0,
		109, 20, 1, 0, 0, 0, 110, 111, 5, 69, 0, 0, 111, 112, 5, 78, 0, 0, 112,
		113, 5, 84, 0, 0, 113, 114, 5, 65, 0, 0, 114, 115, 5, 79, 0, 0, 115, 22,
		1, 0, 0, 0, 116, 117, 5, 83, 0, 0, 117, 118, 5, 69, 0, 0, 118, 119, 5,
		78, 0, 0, 119, 120, 5, 65, 0, 0, 120, 121, 5, 79, 0, 0, 121, 24, 1, 0,
		0, 0, 122, 123, 5, 69, 0, 0, 123, 124, 5, 78, 0, 0, 124, 125, 5, 81, 0,
		0, 125, 126, 5, 85, 0, 0, 126, 127, 5, 65, 0, 0, 127, 128, 5, 78, 0, 0,
		128, 129, 5, 84, 0, 0, 129, 130, 5, 79, 0, 0, 130, 26, 1, 0, 0, 0, 131,
		132, 5, 73, 0, 0, 132, 133, 5, 78, 0, 0, 133, 134, 5, 73, 0, 0, 134, 135,
		5, 67, 0, 0, 135, 136, 5, 73, 0, 0, 136, 137, 5, 79, 0, 0, 137, 28, 1,
		0, 0, 0, 138, 139, 5, 70, 0, 0, 139, 140, 5, 73, 0, 0, 140, 141, 5, 77,
		0, 0, 141, 30, 1, 0, 0, 0, 142, 143, 5, 73, 0, 0, 143, 144, 5, 78, 0, 0,
		144, 145, 5, 84, 0, 0, 145, 146, 5, 69, 0, 0, 146, 147, 5, 73, 0, 0, 147,
		148, 5, 82, 0, 0, 148, 154, 5, 79, 0, 0, 149, 150, 5, 82, 0, 0, 150, 151,
		5, 69, 0, 0, 151, 152, 5, 65, 0, 0, 152, 154, 5, 76, 0, 0, 153, 142, 1,
		0, 0, 0, 153, 149, 1, 0, 0, 0, 154, 32, 1, 0, 0, 0, 155, 157, 2, 48, 57,
		0, 156, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 34, 1, 0, 0, 0, 160, 162, 2, 48, 57, 0, 161, 160,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 171, 1, 0, 0, 0, 165, 167, 5, 46, 0, 0, 166, 168, 2, 48, 57,
		0, 167, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 165, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 36, 1, 0, 0, 0, 173, 177, 7, 0, 0, 0, 174, 176, 7, 1,
		0, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0,
		177, 178, 1, 0, 0, 0, 178, 38, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 185,
		5, 39, 0, 0, 181, 184, 3, 49, 24, 0, 182, 184, 8, 2, 0, 0, 183, 181, 1,
		0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0,
		0, 185, 186, 1, 0, 0, 0, 186, 188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188,
		189, 5, 39, 0, 0, 189, 40, 1, 0, 0, 0, 190, 191, 7, 3, 0, 0, 191, 42, 1,
		0, 0, 0, 192, 193, 7, 4, 0, 0, 193, 44, 1, 0, 0, 0, 194, 204, 5, 62, 0,
		0, 195, 196, 5, 62, 0, 0, 196, 204, 5, 61, 0, 0, 197, 204, 5, 60, 0, 0,
		198, 199, 5, 60, 0, 0, 199, 204, 5, 61, 0, 0, 200, 201, 5, 60, 0, 0, 201,
		204, 5, 62, 0, 0, 202, 204, 5, 61, 0, 0, 203, 194, 1, 0, 0, 0, 203, 195,
		1, 0, 0, 0, 203, 197, 1, 0, 0, 0, 203, 198, 1, 0, 0, 0, 203, 200, 1, 0,
		0, 0, 203, 202, 1, 0, 0, 0, 204, 46, 1, 0, 0, 0, 205, 209, 5, 69, 0, 0,
		206, 207, 5, 79, 0, 0, 207, 209, 5, 85, 0, 0, 208, 205, 1, 0, 0, 0, 208,
		206, 1, 0, 0, 0, 209, 48, 1, 0, 0, 0, 210, 211, 5, 92, 0, 0, 211, 212,
		5, 39, 0, 0, 212, 50, 1, 0, 0, 0, 213, 217, 5, 37, 0, 0, 214, 216, 8, 5,
		0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0,
		217, 218, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220,
		222, 5, 13, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 224, 5, 10, 0, 0, 224, 225, 6, 25, 0, 0, 225, 52, 1, 0,
		0, 0, 226, 227, 7, 6, 0, 0, 227, 228, 6, 26, 1, 0, 228, 54, 1, 0, 0, 0,
		13, 0, 153, 158, 163, 169, 171, 177, 183, 185, 203, 208, 217, 221, 2, 1,
		25, 0, 1, 26, 1,
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

// AlgumaGrammarLexerInit initializes any static state used to implement AlgumaGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAlgumaGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlgumaGrammarLexerInit() {
	staticData := &AlgumaGrammarLexerLexerStaticData
	staticData.once.Do(algumagrammarlexerLexerInit)
}

// NewAlgumaGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAlgumaGrammarLexer(input antlr.CharStream) *AlgumaGrammarLexer {
	AlgumaGrammarLexerInit()
	l := new(AlgumaGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AlgumaGrammarLexerLexerStaticData
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

// AlgumaGrammarLexer tokens.
const (
	AlgumaGrammarLexerT__0       = 1
	AlgumaGrammarLexerT__1       = 2
	AlgumaGrammarLexerT__2       = 3
	AlgumaGrammarLexerT__3       = 4
	AlgumaGrammarLexerT__4       = 5
	AlgumaGrammarLexerT__5       = 6
	AlgumaGrammarLexerT__6       = 7
	AlgumaGrammarLexerT__7       = 8
	AlgumaGrammarLexerT__8       = 9
	AlgumaGrammarLexerT__9       = 10
	AlgumaGrammarLexerT__10      = 11
	AlgumaGrammarLexerT__11      = 12
	AlgumaGrammarLexerT__12      = 13
	AlgumaGrammarLexerT__13      = 14
	AlgumaGrammarLexerT__14      = 15
	AlgumaGrammarLexerTIPO_VAR   = 16
	AlgumaGrammarLexerNUMINT     = 17
	AlgumaGrammarLexerNUMREAL    = 18
	AlgumaGrammarLexerVARIAVEL   = 19
	AlgumaGrammarLexerCADEIA     = 20
	AlgumaGrammarLexerOP_ARIT1   = 21
	AlgumaGrammarLexerOP_ARIT2   = 22
	AlgumaGrammarLexerOP_REL     = 23
	AlgumaGrammarLexerOP_BOOL    = 24
	AlgumaGrammarLexerCOMENTARIO = 25
	AlgumaGrammarLexerWS         = 26
)

func (l *AlgumaGrammarLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 25:
		l.COMENTARIO_Action(localctx, actionIndex)

	case 26:
		l.WS_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *AlgumaGrammarLexer) COMENTARIO_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *AlgumaGrammarLexer) WS_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
