// Code generated from .//parser//AlgumaGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlgumaGrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AlgumaGrammarParser struct {
	*antlr.BaseParser
}

var AlgumaGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algumagrammarParserInit() {
	staticData := &AlgumaGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'DECLARACOES'", "'ALGORITMO'", "'('", "')'", "'ATRIBUIR'",
		"'A'", "'LER'", "'IMPRIMIR'", "'SE'", "'ENTAO'", "'SENAO'", "'ENQUANTO'",
		"'INICIO'", "'FIM'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TIPO_VAR",
		"NUMINT", "NUMREAL", "OP_BOOL", "VARIAVEL", "CADEIA", "OP_ARIT1", "OP_ARIT2",
		"OP_REL", "COMENTARIO", "WS",
	}
	staticData.RuleNames = []string{
		"programa", "listaDeclaracoes", "declaracao", "expressaoAritmetica",
		"termoAritmetico", "fatorAritmetico", "expressaoRelacional", "termoRelacional",
		"listaComandos", "comando", "comandoAtribuicao", "comandoEntrada", "comandoSaida",
		"comandoCondicao", "comandoRepeticao", "subAlgoritmo",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 147, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 41, 8, 1, 11, 1,
		12, 1, 42, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 55, 8, 3, 10, 3, 12, 3, 58, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 5, 4, 66, 8, 4, 10, 4, 12, 4, 69, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 78, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 86,
		8, 6, 10, 6, 12, 6, 89, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 99, 8, 7, 1, 8, 4, 8, 102, 8, 8, 11, 8, 12, 8, 103, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 112, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 137,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 0,
		3, 6, 8, 12, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 0, 1, 1, 0, 20, 21, 145, 0, 32, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 44,
		1, 0, 0, 0, 6, 48, 1, 0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12,
		79, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 101, 1, 0, 0, 0, 18, 111, 1, 0,
		0, 0, 20, 113, 1, 0, 0, 0, 22, 118, 1, 0, 0, 0, 24, 121, 1, 0, 0, 0, 26,
		136, 1, 0, 0, 0, 28, 138, 1, 0, 0, 0, 30, 142, 1, 0, 0, 0, 32, 33, 5, 1,
		0, 0, 33, 34, 5, 2, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 1, 0, 0, 36, 37,
		5, 3, 0, 0, 37, 38, 3, 16, 8, 0, 38, 1, 1, 0, 0, 0, 39, 41, 3, 4, 2, 0,
		40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1,
		0, 0, 0, 43, 3, 1, 0, 0, 0, 44, 45, 5, 20, 0, 0, 45, 46, 5, 1, 0, 0, 46,
		47, 5, 16, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 6, 3, -1, 0, 49, 50, 3, 8,
		4, 0, 50, 56, 1, 0, 0, 0, 51, 52, 10, 2, 0, 0, 52, 53, 5, 22, 0, 0, 53,
		55, 3, 8, 4, 0, 54, 51, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0,
		0, 56, 57, 1, 0, 0, 0, 57, 7, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 6,
		4, -1, 0, 60, 61, 3, 10, 5, 0, 61, 67, 1, 0, 0, 0, 62, 63, 10, 2, 0, 0,
		63, 64, 5, 23, 0, 0, 64, 66, 3, 10, 5, 0, 65, 62, 1, 0, 0, 0, 66, 69, 1,
		0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 9, 1, 0, 0, 0, 69,
		67, 1, 0, 0, 0, 70, 78, 5, 17, 0, 0, 71, 78, 5, 18, 0, 0, 72, 78, 5, 20,
		0, 0, 73, 74, 5, 4, 0, 0, 74, 75, 3, 6, 3, 0, 75, 76, 5, 5, 0, 0, 76, 78,
		1, 0, 0, 0, 77, 70, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0,
		77, 73, 1, 0, 0, 0, 78, 11, 1, 0, 0, 0, 79, 80, 6, 6, -1, 0, 80, 81, 3,
		14, 7, 0, 81, 87, 1, 0, 0, 0, 82, 83, 10, 2, 0, 0, 83, 84, 5, 19, 0, 0,
		84, 86, 3, 14, 7, 0, 85, 82, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1,
		0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 13, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90,
		91, 3, 6, 3, 0, 91, 92, 5, 24, 0, 0, 92, 93, 3, 6, 3, 0, 93, 99, 1, 0,
		0, 0, 94, 95, 5, 4, 0, 0, 95, 96, 3, 12, 6, 0, 96, 97, 5, 5, 0, 0, 97,
		99, 1, 0, 0, 0, 98, 90, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 99, 15, 1, 0, 0,
		0, 100, 102, 3, 18, 9, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 17, 1, 0, 0, 0, 105, 112, 3,
		20, 10, 0, 106, 112, 3, 22, 11, 0, 107, 112, 3, 24, 12, 0, 108, 112, 3,
		26, 13, 0, 109, 112, 3, 28, 14, 0, 110, 112, 3, 30, 15, 0, 111, 105, 1,
		0, 0, 0, 111, 106, 1, 0, 0, 0, 111, 107, 1, 0, 0, 0, 111, 108, 1, 0, 0,
		0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 19, 1, 0, 0, 0, 113,
		114, 5, 6, 0, 0, 114, 115, 3, 6, 3, 0, 115, 116, 5, 7, 0, 0, 116, 117,
		5, 20, 0, 0, 117, 21, 1, 0, 0, 0, 118, 119, 5, 8, 0, 0, 119, 120, 5, 20,
		0, 0, 120, 23, 1, 0, 0, 0, 121, 122, 5, 9, 0, 0, 122, 123, 7, 0, 0, 0,
		123, 25, 1, 0, 0, 0, 124, 125, 5, 10, 0, 0, 125, 126, 3, 12, 6, 0, 126,
		127, 5, 11, 0, 0, 127, 128, 3, 18, 9, 0, 128, 137, 1, 0, 0, 0, 129, 130,
		5, 10, 0, 0, 130, 131, 3, 12, 6, 0, 131, 132, 5, 11, 0, 0, 132, 133, 3,
		18, 9, 0, 133, 134, 5, 12, 0, 0, 134, 135, 3, 18, 9, 0, 135, 137, 1, 0,
		0, 0, 136, 124, 1, 0, 0, 0, 136, 129, 1, 0, 0, 0, 137, 27, 1, 0, 0, 0,
		138, 139, 5, 13, 0, 0, 139, 140, 3, 12, 6, 0, 140, 141, 3, 18, 9, 0, 141,
		29, 1, 0, 0, 0, 142, 143, 5, 14, 0, 0, 143, 144, 3, 16, 8, 0, 144, 145,
		5, 15, 0, 0, 145, 31, 1, 0, 0, 0, 9, 42, 56, 67, 77, 87, 98, 103, 111,
		136,
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

// AlgumaGrammarParserInit initializes any static state used to implement AlgumaGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAlgumaGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlgumaGrammarParserInit() {
	staticData := &AlgumaGrammarParserStaticData
	staticData.once.Do(algumagrammarParserInit)
}

// NewAlgumaGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAlgumaGrammarParser(input antlr.TokenStream) *AlgumaGrammarParser {
	AlgumaGrammarParserInit()
	this := new(AlgumaGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AlgumaGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AlgumaGrammar.g4"

	return this
}

// AlgumaGrammarParser tokens.
const (
	AlgumaGrammarParserEOF        = antlr.TokenEOF
	AlgumaGrammarParserT__0       = 1
	AlgumaGrammarParserT__1       = 2
	AlgumaGrammarParserT__2       = 3
	AlgumaGrammarParserT__3       = 4
	AlgumaGrammarParserT__4       = 5
	AlgumaGrammarParserT__5       = 6
	AlgumaGrammarParserT__6       = 7
	AlgumaGrammarParserT__7       = 8
	AlgumaGrammarParserT__8       = 9
	AlgumaGrammarParserT__9       = 10
	AlgumaGrammarParserT__10      = 11
	AlgumaGrammarParserT__11      = 12
	AlgumaGrammarParserT__12      = 13
	AlgumaGrammarParserT__13      = 14
	AlgumaGrammarParserT__14      = 15
	AlgumaGrammarParserTIPO_VAR   = 16
	AlgumaGrammarParserNUMINT     = 17
	AlgumaGrammarParserNUMREAL    = 18
	AlgumaGrammarParserOP_BOOL    = 19
	AlgumaGrammarParserVARIAVEL   = 20
	AlgumaGrammarParserCADEIA     = 21
	AlgumaGrammarParserOP_ARIT1   = 22
	AlgumaGrammarParserOP_ARIT2   = 23
	AlgumaGrammarParserOP_REL     = 24
	AlgumaGrammarParserCOMENTARIO = 25
	AlgumaGrammarParserWS         = 26
)

// AlgumaGrammarParser rules.
const (
	AlgumaGrammarParserRULE_programa            = 0
	AlgumaGrammarParserRULE_listaDeclaracoes    = 1
	AlgumaGrammarParserRULE_declaracao          = 2
	AlgumaGrammarParserRULE_expressaoAritmetica = 3
	AlgumaGrammarParserRULE_termoAritmetico     = 4
	AlgumaGrammarParserRULE_fatorAritmetico     = 5
	AlgumaGrammarParserRULE_expressaoRelacional = 6
	AlgumaGrammarParserRULE_termoRelacional     = 7
	AlgumaGrammarParserRULE_listaComandos       = 8
	AlgumaGrammarParserRULE_comando             = 9
	AlgumaGrammarParserRULE_comandoAtribuicao   = 10
	AlgumaGrammarParserRULE_comandoEntrada      = 11
	AlgumaGrammarParserRULE_comandoSaida        = 12
	AlgumaGrammarParserRULE_comandoCondicao     = 13
	AlgumaGrammarParserRULE_comandoRepeticao    = 14
	AlgumaGrammarParserRULE_subAlgoritmo        = 15
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaDeclaracoes() IListaDeclaracoesContext
	ListaComandos() IListaComandosContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) ListaDeclaracoes() IListaDeclaracoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaDeclaracoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaDeclaracoesContext)
}

func (s *ProgramaContext) ListaComandos() IListaComandosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaComandosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaComandosContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AlgumaGrammarParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(AlgumaGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(AlgumaGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.ListaDeclaracoes()
	}
	{
		p.SetState(35)
		p.Match(AlgumaGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(AlgumaGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.ListaComandos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaDeclaracoesContext is an interface to support dynamic dispatch.
type IListaDeclaracoesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracao() []IDeclaracaoContext
	Declaracao(i int) IDeclaracaoContext

	// IsListaDeclaracoesContext differentiates from other interfaces.
	IsListaDeclaracoesContext()
}

type ListaDeclaracoesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaDeclaracoesContext() *ListaDeclaracoesContext {
	var p = new(ListaDeclaracoesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_listaDeclaracoes
	return p
}

func InitEmptyListaDeclaracoesContext(p *ListaDeclaracoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_listaDeclaracoes
}

func (*ListaDeclaracoesContext) IsListaDeclaracoesContext() {}

func NewListaDeclaracoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaDeclaracoesContext {
	var p = new(ListaDeclaracoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_listaDeclaracoes

	return p
}

func (s *ListaDeclaracoesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaDeclaracoesContext) AllDeclaracao() []IDeclaracaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracaoContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracaoContext); ok {
			tst[i] = t.(IDeclaracaoContext)
			i++
		}
	}

	return tst
}

func (s *ListaDeclaracoesContext) Declaracao(i int) IDeclaracaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracaoContext)
}

func (s *ListaDeclaracoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaDeclaracoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaDeclaracoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterListaDeclaracoes(s)
	}
}

func (s *ListaDeclaracoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitListaDeclaracoes(s)
	}
}

func (s *ListaDeclaracoesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitListaDeclaracoes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ListaDeclaracoes() (localctx IListaDeclaracoesContext) {
	localctx = NewListaDeclaracoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AlgumaGrammarParserRULE_listaDeclaracoes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AlgumaGrammarParserVARIAVEL {
		{
			p.SetState(39)
			p.Declaracao()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracaoContext is an interface to support dynamic dispatch.
type IDeclaracaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIAVEL() antlr.TerminalNode
	TIPO_VAR() antlr.TerminalNode

	// IsDeclaracaoContext differentiates from other interfaces.
	IsDeclaracaoContext()
}

type DeclaracaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracaoContext() *DeclaracaoContext {
	var p = new(DeclaracaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_declaracao
	return p
}

func InitEmptyDeclaracaoContext(p *DeclaracaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_declaracao
}

func (*DeclaracaoContext) IsDeclaracaoContext() {}

func NewDeclaracaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracaoContext {
	var p = new(DeclaracaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_declaracao

	return p
}

func (s *DeclaracaoContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracaoContext) VARIAVEL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserVARIAVEL, 0)
}

func (s *DeclaracaoContext) TIPO_VAR() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserTIPO_VAR, 0)
}

func (s *DeclaracaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterDeclaracao(s)
	}
}

func (s *DeclaracaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitDeclaracao(s)
	}
}

func (s *DeclaracaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitDeclaracao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) Declaracao() (localctx IDeclaracaoContext) {
	localctx = NewDeclaracaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AlgumaGrammarParserRULE_declaracao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(AlgumaGrammarParserVARIAVEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(AlgumaGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(AlgumaGrammarParserTIPO_VAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressaoAritmeticaContext is an interface to support dynamic dispatch.
type IExpressaoAritmeticaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TermoAritmetico() ITermoAritmeticoContext
	ExpressaoAritmetica() IExpressaoAritmeticaContext
	OP_ARIT1() antlr.TerminalNode

	// IsExpressaoAritmeticaContext differentiates from other interfaces.
	IsExpressaoAritmeticaContext()
}

type ExpressaoAritmeticaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressaoAritmeticaContext() *ExpressaoAritmeticaContext {
	var p = new(ExpressaoAritmeticaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoAritmetica
	return p
}

func InitEmptyExpressaoAritmeticaContext(p *ExpressaoAritmeticaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoAritmetica
}

func (*ExpressaoAritmeticaContext) IsExpressaoAritmeticaContext() {}

func NewExpressaoAritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoAritmeticaContext {
	var p = new(ExpressaoAritmeticaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoAritmetica

	return p
}

func (s *ExpressaoAritmeticaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressaoAritmeticaContext) TermoAritmetico() ITermoAritmeticoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoAritmeticoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoAritmeticoContext)
}

func (s *ExpressaoAritmeticaContext) ExpressaoAritmetica() IExpressaoAritmeticaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoAritmeticaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoAritmeticaContext)
}

func (s *ExpressaoAritmeticaContext) OP_ARIT1() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserOP_ARIT1, 0)
}

func (s *ExpressaoAritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressaoAritmeticaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressaoAritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterExpressaoAritmetica(s)
	}
}

func (s *ExpressaoAritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitExpressaoAritmetica(s)
	}
}

func (s *ExpressaoAritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitExpressaoAritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ExpressaoAritmetica() (localctx IExpressaoAritmeticaContext) {
	return p.expressaoAritmetica(0)
}

func (p *AlgumaGrammarParser) expressaoAritmetica(_p int) (localctx IExpressaoAritmeticaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressaoAritmeticaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressaoAritmeticaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, AlgumaGrammarParserRULE_expressaoAritmetica, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.termoAritmetico(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressaoAritmeticaContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AlgumaGrammarParserRULE_expressaoAritmetica)
			p.SetState(51)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(52)
				p.Match(AlgumaGrammarParserOP_ARIT1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(53)
				p.termoAritmetico(0)
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoAritmeticoContext is an interface to support dynamic dispatch.
type ITermoAritmeticoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FatorAritmetico() IFatorAritmeticoContext
	TermoAritmetico() ITermoAritmeticoContext
	OP_ARIT2() antlr.TerminalNode

	// IsTermoAritmeticoContext differentiates from other interfaces.
	IsTermoAritmeticoContext()
}

type TermoAritmeticoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoAritmeticoContext() *TermoAritmeticoContext {
	var p = new(TermoAritmeticoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_termoAritmetico
	return p
}

func InitEmptyTermoAritmeticoContext(p *TermoAritmeticoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_termoAritmetico
}

func (*TermoAritmeticoContext) IsTermoAritmeticoContext() {}

func NewTermoAritmeticoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoAritmeticoContext {
	var p = new(TermoAritmeticoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_termoAritmetico

	return p
}

func (s *TermoAritmeticoContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoAritmeticoContext) FatorAritmetico() IFatorAritmeticoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorAritmeticoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorAritmeticoContext)
}

func (s *TermoAritmeticoContext) TermoAritmetico() ITermoAritmeticoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoAritmeticoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoAritmeticoContext)
}

func (s *TermoAritmeticoContext) OP_ARIT2() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserOP_ARIT2, 0)
}

func (s *TermoAritmeticoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoAritmeticoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoAritmeticoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterTermoAritmetico(s)
	}
}

func (s *TermoAritmeticoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitTermoAritmetico(s)
	}
}

func (s *TermoAritmeticoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitTermoAritmetico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) TermoAritmetico() (localctx ITermoAritmeticoContext) {
	return p.termoAritmetico(0)
}

func (p *AlgumaGrammarParser) termoAritmetico(_p int) (localctx ITermoAritmeticoContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTermoAritmeticoContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermoAritmeticoContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, AlgumaGrammarParserRULE_termoAritmetico, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.FatorAritmetico()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermoAritmeticoContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AlgumaGrammarParserRULE_termoAritmetico)
			p.SetState(62)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(63)
				p.Match(AlgumaGrammarParserOP_ARIT2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(64)
				p.FatorAritmetico()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFatorAritmeticoContext is an interface to support dynamic dispatch.
type IFatorAritmeticoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMINT() antlr.TerminalNode
	NUMREAL() antlr.TerminalNode
	VARIAVEL() antlr.TerminalNode
	ExpressaoAritmetica() IExpressaoAritmeticaContext

	// IsFatorAritmeticoContext differentiates from other interfaces.
	IsFatorAritmeticoContext()
}

type FatorAritmeticoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatorAritmeticoContext() *FatorAritmeticoContext {
	var p = new(FatorAritmeticoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_fatorAritmetico
	return p
}

func InitEmptyFatorAritmeticoContext(p *FatorAritmeticoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_fatorAritmetico
}

func (*FatorAritmeticoContext) IsFatorAritmeticoContext() {}

func NewFatorAritmeticoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatorAritmeticoContext {
	var p = new(FatorAritmeticoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_fatorAritmetico

	return p
}

func (s *FatorAritmeticoContext) GetParser() antlr.Parser { return s.parser }

func (s *FatorAritmeticoContext) NUMINT() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserNUMINT, 0)
}

func (s *FatorAritmeticoContext) NUMREAL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserNUMREAL, 0)
}

func (s *FatorAritmeticoContext) VARIAVEL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserVARIAVEL, 0)
}

func (s *FatorAritmeticoContext) ExpressaoAritmetica() IExpressaoAritmeticaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoAritmeticaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoAritmeticaContext)
}

func (s *FatorAritmeticoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatorAritmeticoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FatorAritmeticoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterFatorAritmetico(s)
	}
}

func (s *FatorAritmeticoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitFatorAritmetico(s)
	}
}

func (s *FatorAritmeticoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitFatorAritmetico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) FatorAritmetico() (localctx IFatorAritmeticoContext) {
	localctx = NewFatorAritmeticoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AlgumaGrammarParserRULE_fatorAritmetico)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlgumaGrammarParserNUMINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(AlgumaGrammarParserNUMINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlgumaGrammarParserNUMREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(AlgumaGrammarParserNUMREAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlgumaGrammarParserVARIAVEL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(AlgumaGrammarParserVARIAVEL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlgumaGrammarParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Match(AlgumaGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.expressaoAritmetica(0)
		}
		{
			p.SetState(75)
			p.Match(AlgumaGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressaoRelacionalContext is an interface to support dynamic dispatch.
type IExpressaoRelacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TermoRelacional() ITermoRelacionalContext
	ExpressaoRelacional() IExpressaoRelacionalContext
	OP_BOOL() antlr.TerminalNode

	// IsExpressaoRelacionalContext differentiates from other interfaces.
	IsExpressaoRelacionalContext()
}

type ExpressaoRelacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressaoRelacionalContext() *ExpressaoRelacionalContext {
	var p = new(ExpressaoRelacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoRelacional
	return p
}

func InitEmptyExpressaoRelacionalContext(p *ExpressaoRelacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoRelacional
}

func (*ExpressaoRelacionalContext) IsExpressaoRelacionalContext() {}

func NewExpressaoRelacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoRelacionalContext {
	var p = new(ExpressaoRelacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_expressaoRelacional

	return p
}

func (s *ExpressaoRelacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressaoRelacionalContext) TermoRelacional() ITermoRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoRelacionalContext)
}

func (s *ExpressaoRelacionalContext) ExpressaoRelacional() IExpressaoRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoRelacionalContext)
}

func (s *ExpressaoRelacionalContext) OP_BOOL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserOP_BOOL, 0)
}

func (s *ExpressaoRelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressaoRelacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressaoRelacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterExpressaoRelacional(s)
	}
}

func (s *ExpressaoRelacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitExpressaoRelacional(s)
	}
}

func (s *ExpressaoRelacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitExpressaoRelacional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ExpressaoRelacional() (localctx IExpressaoRelacionalContext) {
	return p.expressaoRelacional(0)
}

func (p *AlgumaGrammarParser) expressaoRelacional(_p int) (localctx IExpressaoRelacionalContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressaoRelacionalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressaoRelacionalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, AlgumaGrammarParserRULE_expressaoRelacional, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.TermoRelacional()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressaoRelacionalContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AlgumaGrammarParserRULE_expressaoRelacional)
			p.SetState(82)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(83)
				p.Match(AlgumaGrammarParserOP_BOOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(84)
				p.TermoRelacional()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoRelacionalContext is an interface to support dynamic dispatch.
type ITermoRelacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressaoAritmetica() []IExpressaoAritmeticaContext
	ExpressaoAritmetica(i int) IExpressaoAritmeticaContext
	OP_REL() antlr.TerminalNode
	ExpressaoRelacional() IExpressaoRelacionalContext

	// IsTermoRelacionalContext differentiates from other interfaces.
	IsTermoRelacionalContext()
}

type TermoRelacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoRelacionalContext() *TermoRelacionalContext {
	var p = new(TermoRelacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_termoRelacional
	return p
}

func InitEmptyTermoRelacionalContext(p *TermoRelacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_termoRelacional
}

func (*TermoRelacionalContext) IsTermoRelacionalContext() {}

func NewTermoRelacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoRelacionalContext {
	var p = new(TermoRelacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_termoRelacional

	return p
}

func (s *TermoRelacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoRelacionalContext) AllExpressaoAritmetica() []IExpressaoAritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoAritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoAritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoAritmeticaContext); ok {
			tst[i] = t.(IExpressaoAritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *TermoRelacionalContext) ExpressaoAritmetica(i int) IExpressaoAritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoAritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoAritmeticaContext)
}

func (s *TermoRelacionalContext) OP_REL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserOP_REL, 0)
}

func (s *TermoRelacionalContext) ExpressaoRelacional() IExpressaoRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoRelacionalContext)
}

func (s *TermoRelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoRelacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoRelacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterTermoRelacional(s)
	}
}

func (s *TermoRelacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitTermoRelacional(s)
	}
}

func (s *TermoRelacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitTermoRelacional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) TermoRelacional() (localctx ITermoRelacionalContext) {
	localctx = NewTermoRelacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AlgumaGrammarParserRULE_termoRelacional)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.expressaoAritmetica(0)
		}
		{
			p.SetState(91)
			p.Match(AlgumaGrammarParserOP_REL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.expressaoAritmetica(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(AlgumaGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.expressaoRelacional(0)
		}
		{
			p.SetState(96)
			p.Match(AlgumaGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaComandosContext is an interface to support dynamic dispatch.
type IListaComandosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComando() []IComandoContext
	Comando(i int) IComandoContext

	// IsListaComandosContext differentiates from other interfaces.
	IsListaComandosContext()
}

type ListaComandosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaComandosContext() *ListaComandosContext {
	var p = new(ListaComandosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_listaComandos
	return p
}

func InitEmptyListaComandosContext(p *ListaComandosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_listaComandos
}

func (*ListaComandosContext) IsListaComandosContext() {}

func NewListaComandosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaComandosContext {
	var p = new(ListaComandosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_listaComandos

	return p
}

func (s *ListaComandosContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaComandosContext) AllComando() []IComandoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComandoContext); ok {
			len++
		}
	}

	tst := make([]IComandoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComandoContext); ok {
			tst[i] = t.(IComandoContext)
			i++
		}
	}

	return tst
}

func (s *ListaComandosContext) Comando(i int) IComandoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *ListaComandosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaComandosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaComandosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterListaComandos(s)
	}
}

func (s *ListaComandosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitListaComandos(s)
	}
}

func (s *ListaComandosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitListaComandos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ListaComandos() (localctx IListaComandosContext) {
	localctx = NewListaComandosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AlgumaGrammarParserRULE_listaComandos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26432) != 0) {
		{
			p.SetState(100)
			p.Comando()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComandoAtribuicao() IComandoAtribuicaoContext
	ComandoEntrada() IComandoEntradaContext
	ComandoSaida() IComandoSaidaContext
	ComandoCondicao() IComandoCondicaoContext
	ComandoRepeticao() IComandoRepeticaoContext
	SubAlgoritmo() ISubAlgoritmoContext

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoContext() *ComandoContext {
	var p = new(ComandoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comando
	return p
}

func InitEmptyComandoContext(p *ComandoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comando
}

func (*ComandoContext) IsComandoContext() {}

func NewComandoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoContext {
	var p = new(ComandoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comando

	return p
}

func (s *ComandoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoContext) ComandoAtribuicao() IComandoAtribuicaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoAtribuicaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoAtribuicaoContext)
}

func (s *ComandoContext) ComandoEntrada() IComandoEntradaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoEntradaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoEntradaContext)
}

func (s *ComandoContext) ComandoSaida() IComandoSaidaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoSaidaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoSaidaContext)
}

func (s *ComandoContext) ComandoCondicao() IComandoCondicaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoCondicaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoCondicaoContext)
}

func (s *ComandoContext) ComandoRepeticao() IComandoRepeticaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoRepeticaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoRepeticaoContext)
}

func (s *ComandoContext) SubAlgoritmo() ISubAlgoritmoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAlgoritmoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAlgoritmoContext)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComando(s)
	}
}

func (s *ComandoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComando(s)
	}
}

func (s *ComandoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComando(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) Comando() (localctx IComandoContext) {
	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AlgumaGrammarParserRULE_comando)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlgumaGrammarParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.ComandoAtribuicao()
		}

	case AlgumaGrammarParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.ComandoEntrada()
		}

	case AlgumaGrammarParserT__8:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.ComandoSaida()
		}

	case AlgumaGrammarParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.ComandoCondicao()
		}

	case AlgumaGrammarParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.ComandoRepeticao()
		}

	case AlgumaGrammarParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.SubAlgoritmo()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoAtribuicaoContext is an interface to support dynamic dispatch.
type IComandoAtribuicaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressaoAritmetica() IExpressaoAritmeticaContext
	VARIAVEL() antlr.TerminalNode

	// IsComandoAtribuicaoContext differentiates from other interfaces.
	IsComandoAtribuicaoContext()
}

type ComandoAtribuicaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoAtribuicaoContext() *ComandoAtribuicaoContext {
	var p = new(ComandoAtribuicaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoAtribuicao
	return p
}

func InitEmptyComandoAtribuicaoContext(p *ComandoAtribuicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoAtribuicao
}

func (*ComandoAtribuicaoContext) IsComandoAtribuicaoContext() {}

func NewComandoAtribuicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoAtribuicaoContext {
	var p = new(ComandoAtribuicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comandoAtribuicao

	return p
}

func (s *ComandoAtribuicaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoAtribuicaoContext) ExpressaoAritmetica() IExpressaoAritmeticaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoAritmeticaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoAritmeticaContext)
}

func (s *ComandoAtribuicaoContext) VARIAVEL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserVARIAVEL, 0)
}

func (s *ComandoAtribuicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoAtribuicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoAtribuicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComandoAtribuicao(s)
	}
}

func (s *ComandoAtribuicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComandoAtribuicao(s)
	}
}

func (s *ComandoAtribuicaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComandoAtribuicao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ComandoAtribuicao() (localctx IComandoAtribuicaoContext) {
	localctx = NewComandoAtribuicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AlgumaGrammarParserRULE_comandoAtribuicao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(AlgumaGrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.expressaoAritmetica(0)
	}
	{
		p.SetState(115)
		p.Match(AlgumaGrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(AlgumaGrammarParserVARIAVEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoEntradaContext is an interface to support dynamic dispatch.
type IComandoEntradaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIAVEL() antlr.TerminalNode

	// IsComandoEntradaContext differentiates from other interfaces.
	IsComandoEntradaContext()
}

type ComandoEntradaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoEntradaContext() *ComandoEntradaContext {
	var p = new(ComandoEntradaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoEntrada
	return p
}

func InitEmptyComandoEntradaContext(p *ComandoEntradaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoEntrada
}

func (*ComandoEntradaContext) IsComandoEntradaContext() {}

func NewComandoEntradaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoEntradaContext {
	var p = new(ComandoEntradaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comandoEntrada

	return p
}

func (s *ComandoEntradaContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoEntradaContext) VARIAVEL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserVARIAVEL, 0)
}

func (s *ComandoEntradaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoEntradaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoEntradaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComandoEntrada(s)
	}
}

func (s *ComandoEntradaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComandoEntrada(s)
	}
}

func (s *ComandoEntradaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComandoEntrada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ComandoEntrada() (localctx IComandoEntradaContext) {
	localctx = NewComandoEntradaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AlgumaGrammarParserRULE_comandoEntrada)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(AlgumaGrammarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(AlgumaGrammarParserVARIAVEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoSaidaContext is an interface to support dynamic dispatch.
type IComandoSaidaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIAVEL() antlr.TerminalNode
	CADEIA() antlr.TerminalNode

	// IsComandoSaidaContext differentiates from other interfaces.
	IsComandoSaidaContext()
}

type ComandoSaidaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoSaidaContext() *ComandoSaidaContext {
	var p = new(ComandoSaidaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoSaida
	return p
}

func InitEmptyComandoSaidaContext(p *ComandoSaidaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoSaida
}

func (*ComandoSaidaContext) IsComandoSaidaContext() {}

func NewComandoSaidaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoSaidaContext {
	var p = new(ComandoSaidaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comandoSaida

	return p
}

func (s *ComandoSaidaContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoSaidaContext) VARIAVEL() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserVARIAVEL, 0)
}

func (s *ComandoSaidaContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(AlgumaGrammarParserCADEIA, 0)
}

func (s *ComandoSaidaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoSaidaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoSaidaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComandoSaida(s)
	}
}

func (s *ComandoSaidaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComandoSaida(s)
	}
}

func (s *ComandoSaidaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComandoSaida(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ComandoSaida() (localctx IComandoSaidaContext) {
	localctx = NewComandoSaidaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AlgumaGrammarParserRULE_comandoSaida)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(AlgumaGrammarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AlgumaGrammarParserVARIAVEL || _la == AlgumaGrammarParserCADEIA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoCondicaoContext is an interface to support dynamic dispatch.
type IComandoCondicaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressaoRelacional() IExpressaoRelacionalContext
	AllComando() []IComandoContext
	Comando(i int) IComandoContext

	// IsComandoCondicaoContext differentiates from other interfaces.
	IsComandoCondicaoContext()
}

type ComandoCondicaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoCondicaoContext() *ComandoCondicaoContext {
	var p = new(ComandoCondicaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoCondicao
	return p
}

func InitEmptyComandoCondicaoContext(p *ComandoCondicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoCondicao
}

func (*ComandoCondicaoContext) IsComandoCondicaoContext() {}

func NewComandoCondicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoCondicaoContext {
	var p = new(ComandoCondicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comandoCondicao

	return p
}

func (s *ComandoCondicaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoCondicaoContext) ExpressaoRelacional() IExpressaoRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoRelacionalContext)
}

func (s *ComandoCondicaoContext) AllComando() []IComandoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComandoContext); ok {
			len++
		}
	}

	tst := make([]IComandoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComandoContext); ok {
			tst[i] = t.(IComandoContext)
			i++
		}
	}

	return tst
}

func (s *ComandoCondicaoContext) Comando(i int) IComandoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *ComandoCondicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoCondicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoCondicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComandoCondicao(s)
	}
}

func (s *ComandoCondicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComandoCondicao(s)
	}
}

func (s *ComandoCondicaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComandoCondicao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ComandoCondicao() (localctx IComandoCondicaoContext) {
	localctx = NewComandoCondicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AlgumaGrammarParserRULE_comandoCondicao)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(AlgumaGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expressaoRelacional(0)
		}
		{
			p.SetState(126)
			p.Match(AlgumaGrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Comando()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(AlgumaGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expressaoRelacional(0)
		}
		{
			p.SetState(131)
			p.Match(AlgumaGrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Comando()
		}
		{
			p.SetState(133)
			p.Match(AlgumaGrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Comando()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoRepeticaoContext is an interface to support dynamic dispatch.
type IComandoRepeticaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressaoRelacional() IExpressaoRelacionalContext
	Comando() IComandoContext

	// IsComandoRepeticaoContext differentiates from other interfaces.
	IsComandoRepeticaoContext()
}

type ComandoRepeticaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoRepeticaoContext() *ComandoRepeticaoContext {
	var p = new(ComandoRepeticaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoRepeticao
	return p
}

func InitEmptyComandoRepeticaoContext(p *ComandoRepeticaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_comandoRepeticao
}

func (*ComandoRepeticaoContext) IsComandoRepeticaoContext() {}

func NewComandoRepeticaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoRepeticaoContext {
	var p = new(ComandoRepeticaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_comandoRepeticao

	return p
}

func (s *ComandoRepeticaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoRepeticaoContext) ExpressaoRelacional() IExpressaoRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoRelacionalContext)
}

func (s *ComandoRepeticaoContext) Comando() IComandoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *ComandoRepeticaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoRepeticaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoRepeticaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterComandoRepeticao(s)
	}
}

func (s *ComandoRepeticaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitComandoRepeticao(s)
	}
}

func (s *ComandoRepeticaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitComandoRepeticao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) ComandoRepeticao() (localctx IComandoRepeticaoContext) {
	localctx = NewComandoRepeticaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AlgumaGrammarParserRULE_comandoRepeticao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(AlgumaGrammarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.expressaoRelacional(0)
	}
	{
		p.SetState(140)
		p.Comando()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubAlgoritmoContext is an interface to support dynamic dispatch.
type ISubAlgoritmoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaComandos() IListaComandosContext

	// IsSubAlgoritmoContext differentiates from other interfaces.
	IsSubAlgoritmoContext()
}

type SubAlgoritmoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAlgoritmoContext() *SubAlgoritmoContext {
	var p = new(SubAlgoritmoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_subAlgoritmo
	return p
}

func InitEmptySubAlgoritmoContext(p *SubAlgoritmoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlgumaGrammarParserRULE_subAlgoritmo
}

func (*SubAlgoritmoContext) IsSubAlgoritmoContext() {}

func NewSubAlgoritmoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAlgoritmoContext {
	var p = new(SubAlgoritmoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgumaGrammarParserRULE_subAlgoritmo

	return p
}

func (s *SubAlgoritmoContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAlgoritmoContext) ListaComandos() IListaComandosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaComandosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaComandosContext)
}

func (s *SubAlgoritmoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAlgoritmoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAlgoritmoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.EnterSubAlgoritmo(s)
	}
}

func (s *SubAlgoritmoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlgumaGrammarListener); ok {
		listenerT.ExitSubAlgoritmo(s)
	}
}

func (s *SubAlgoritmoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlgumaGrammarVisitor:
		return t.VisitSubAlgoritmo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlgumaGrammarParser) SubAlgoritmo() (localctx ISubAlgoritmoContext) {
	localctx = NewSubAlgoritmoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AlgumaGrammarParserRULE_subAlgoritmo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(AlgumaGrammarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.ListaComandos()
	}
	{
		p.SetState(144)
		p.Match(AlgumaGrammarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *AlgumaGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressaoAritmeticaContext = nil
		if localctx != nil {
			t = localctx.(*ExpressaoAritmeticaContext)
		}
		return p.ExpressaoAritmetica_Sempred(t, predIndex)

	case 4:
		var t *TermoAritmeticoContext = nil
		if localctx != nil {
			t = localctx.(*TermoAritmeticoContext)
		}
		return p.TermoAritmetico_Sempred(t, predIndex)

	case 6:
		var t *ExpressaoRelacionalContext = nil
		if localctx != nil {
			t = localctx.(*ExpressaoRelacionalContext)
		}
		return p.ExpressaoRelacional_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlgumaGrammarParser) ExpressaoAritmetica_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AlgumaGrammarParser) TermoAritmetico_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AlgumaGrammarParser) ExpressaoRelacional_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
