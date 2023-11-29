// Code generated from ./parser/AlgumaGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlgumaGrammar

import (
	"github.com/antlr4-go/antlr"
)

// BaseAlgumaGrammarListener is a complete listener for a parse tree produced by AlgumaGrammarParser.
type BaseAlgumaGrammarListener struct{}

var _ AlgumaGrammarListener = &BaseAlgumaGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlgumaGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlgumaGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlgumaGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlgumaGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseAlgumaGrammarListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseAlgumaGrammarListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterListaDeclaracoes is called when production listaDeclaracoes is entered.
func (s *BaseAlgumaGrammarListener) EnterListaDeclaracoes(ctx *ListaDeclaracoesContext) {}

// ExitListaDeclaracoes is called when production listaDeclaracoes is exited.
func (s *BaseAlgumaGrammarListener) ExitListaDeclaracoes(ctx *ListaDeclaracoesContext) {}

// EnterDeclaracao is called when production declaracao is entered.
func (s *BaseAlgumaGrammarListener) EnterDeclaracao(ctx *DeclaracaoContext) {}

// ExitDeclaracao is called when production declaracao is exited.
func (s *BaseAlgumaGrammarListener) ExitDeclaracao(ctx *DeclaracaoContext) {}

// EnterExpressaoAritmetica is called when production expressaoAritmetica is entered.
func (s *BaseAlgumaGrammarListener) EnterExpressaoAritmetica(ctx *ExpressaoAritmeticaContext) {}

// ExitExpressaoAritmetica is called when production expressaoAritmetica is exited.
func (s *BaseAlgumaGrammarListener) ExitExpressaoAritmetica(ctx *ExpressaoAritmeticaContext) {}

// EnterTermoAritmetico is called when production termoAritmetico is entered.
func (s *BaseAlgumaGrammarListener) EnterTermoAritmetico(ctx *TermoAritmeticoContext) {}

// ExitTermoAritmetico is called when production termoAritmetico is exited.
func (s *BaseAlgumaGrammarListener) ExitTermoAritmetico(ctx *TermoAritmeticoContext) {}

// EnterFatorAritmetico is called when production fatorAritmetico is entered.
func (s *BaseAlgumaGrammarListener) EnterFatorAritmetico(ctx *FatorAritmeticoContext) {}

// ExitFatorAritmetico is called when production fatorAritmetico is exited.
func (s *BaseAlgumaGrammarListener) ExitFatorAritmetico(ctx *FatorAritmeticoContext) {}

// EnterExpressaoRelacional is called when production expressaoRelacional is entered.
func (s *BaseAlgumaGrammarListener) EnterExpressaoRelacional(ctx *ExpressaoRelacionalContext) {}

// ExitExpressaoRelacional is called when production expressaoRelacional is exited.
func (s *BaseAlgumaGrammarListener) ExitExpressaoRelacional(ctx *ExpressaoRelacionalContext) {}

// EnterTermoRelacional is called when production termoRelacional is entered.
func (s *BaseAlgumaGrammarListener) EnterTermoRelacional(ctx *TermoRelacionalContext) {}

// ExitTermoRelacional is called when production termoRelacional is exited.
func (s *BaseAlgumaGrammarListener) ExitTermoRelacional(ctx *TermoRelacionalContext) {}

// EnterListaComandos is called when production listaComandos is entered.
func (s *BaseAlgumaGrammarListener) EnterListaComandos(ctx *ListaComandosContext) {}

// ExitListaComandos is called when production listaComandos is exited.
func (s *BaseAlgumaGrammarListener) ExitListaComandos(ctx *ListaComandosContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseAlgumaGrammarListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseAlgumaGrammarListener) ExitComando(ctx *ComandoContext) {}

// EnterComandoAtribuicao is called when production comandoAtribuicao is entered.
func (s *BaseAlgumaGrammarListener) EnterComandoAtribuicao(ctx *ComandoAtribuicaoContext) {}

// ExitComandoAtribuicao is called when production comandoAtribuicao is exited.
func (s *BaseAlgumaGrammarListener) ExitComandoAtribuicao(ctx *ComandoAtribuicaoContext) {}

// EnterComandoEntrada is called when production comandoEntrada is entered.
func (s *BaseAlgumaGrammarListener) EnterComandoEntrada(ctx *ComandoEntradaContext) {}

// ExitComandoEntrada is called when production comandoEntrada is exited.
func (s *BaseAlgumaGrammarListener) ExitComandoEntrada(ctx *ComandoEntradaContext) {}

// EnterComandoSaida is called when production comandoSaida is entered.
func (s *BaseAlgumaGrammarListener) EnterComandoSaida(ctx *ComandoSaidaContext) {}

// ExitComandoSaida is called when production comandoSaida is exited.
func (s *BaseAlgumaGrammarListener) ExitComandoSaida(ctx *ComandoSaidaContext) {}

// EnterComandoCondicao is called when production comandoCondicao is entered.
func (s *BaseAlgumaGrammarListener) EnterComandoCondicao(ctx *ComandoCondicaoContext) {}

// ExitComandoCondicao is called when production comandoCondicao is exited.
func (s *BaseAlgumaGrammarListener) ExitComandoCondicao(ctx *ComandoCondicaoContext) {}

// EnterComandoRepeticao is called when production comandoRepeticao is entered.
func (s *BaseAlgumaGrammarListener) EnterComandoRepeticao(ctx *ComandoRepeticaoContext) {}

// ExitComandoRepeticao is called when production comandoRepeticao is exited.
func (s *BaseAlgumaGrammarListener) ExitComandoRepeticao(ctx *ComandoRepeticaoContext) {}

// EnterSubAlgoritmo is called when production subAlgoritmo is entered.
func (s *BaseAlgumaGrammarListener) EnterSubAlgoritmo(ctx *SubAlgoritmoContext) {}

// ExitSubAlgoritmo is called when production subAlgoritmo is exited.
func (s *BaseAlgumaGrammarListener) ExitSubAlgoritmo(ctx *SubAlgoritmoContext) {}
