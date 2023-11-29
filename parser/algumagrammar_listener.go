// Code generated from ./parser/AlgumaGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlgumaGrammar

import (
	"github.com/antlr4-go/antlr"
)

// AlgumaGrammarListener is a complete listener for a parse tree produced by AlgumaGrammarParser.
type AlgumaGrammarListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterListaDeclaracoes is called when entering the listaDeclaracoes production.
	EnterListaDeclaracoes(c *ListaDeclaracoesContext)

	// EnterDeclaracao is called when entering the declaracao production.
	EnterDeclaracao(c *DeclaracaoContext)

	// EnterExpressaoAritmetica is called when entering the expressaoAritmetica production.
	EnterExpressaoAritmetica(c *ExpressaoAritmeticaContext)

	// EnterTermoAritmetico is called when entering the termoAritmetico production.
	EnterTermoAritmetico(c *TermoAritmeticoContext)

	// EnterFatorAritmetico is called when entering the fatorAritmetico production.
	EnterFatorAritmetico(c *FatorAritmeticoContext)

	// EnterExpressaoRelacional is called when entering the expressaoRelacional production.
	EnterExpressaoRelacional(c *ExpressaoRelacionalContext)

	// EnterTermoRelacional is called when entering the termoRelacional production.
	EnterTermoRelacional(c *TermoRelacionalContext)

	// EnterListaComandos is called when entering the listaComandos production.
	EnterListaComandos(c *ListaComandosContext)

	// EnterComando is called when entering the comando production.
	EnterComando(c *ComandoContext)

	// EnterComandoAtribuicao is called when entering the comandoAtribuicao production.
	EnterComandoAtribuicao(c *ComandoAtribuicaoContext)

	// EnterComandoEntrada is called when entering the comandoEntrada production.
	EnterComandoEntrada(c *ComandoEntradaContext)

	// EnterComandoSaida is called when entering the comandoSaida production.
	EnterComandoSaida(c *ComandoSaidaContext)

	// EnterComandoCondicao is called when entering the comandoCondicao production.
	EnterComandoCondicao(c *ComandoCondicaoContext)

	// EnterComandoRepeticao is called when entering the comandoRepeticao production.
	EnterComandoRepeticao(c *ComandoRepeticaoContext)

	// EnterSubAlgoritmo is called when entering the subAlgoritmo production.
	EnterSubAlgoritmo(c *SubAlgoritmoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitListaDeclaracoes is called when exiting the listaDeclaracoes production.
	ExitListaDeclaracoes(c *ListaDeclaracoesContext)

	// ExitDeclaracao is called when exiting the declaracao production.
	ExitDeclaracao(c *DeclaracaoContext)

	// ExitExpressaoAritmetica is called when exiting the expressaoAritmetica production.
	ExitExpressaoAritmetica(c *ExpressaoAritmeticaContext)

	// ExitTermoAritmetico is called when exiting the termoAritmetico production.
	ExitTermoAritmetico(c *TermoAritmeticoContext)

	// ExitFatorAritmetico is called when exiting the fatorAritmetico production.
	ExitFatorAritmetico(c *FatorAritmeticoContext)

	// ExitExpressaoRelacional is called when exiting the expressaoRelacional production.
	ExitExpressaoRelacional(c *ExpressaoRelacionalContext)

	// ExitTermoRelacional is called when exiting the termoRelacional production.
	ExitTermoRelacional(c *TermoRelacionalContext)

	// ExitListaComandos is called when exiting the listaComandos production.
	ExitListaComandos(c *ListaComandosContext)

	// ExitComando is called when exiting the comando production.
	ExitComando(c *ComandoContext)

	// ExitComandoAtribuicao is called when exiting the comandoAtribuicao production.
	ExitComandoAtribuicao(c *ComandoAtribuicaoContext)

	// ExitComandoEntrada is called when exiting the comandoEntrada production.
	ExitComandoEntrada(c *ComandoEntradaContext)

	// ExitComandoSaida is called when exiting the comandoSaida production.
	ExitComandoSaida(c *ComandoSaidaContext)

	// ExitComandoCondicao is called when exiting the comandoCondicao production.
	ExitComandoCondicao(c *ComandoCondicaoContext)

	// ExitComandoRepeticao is called when exiting the comandoRepeticao production.
	ExitComandoRepeticao(c *ComandoRepeticaoContext)

	// ExitSubAlgoritmo is called when exiting the subAlgoritmo production.
	ExitSubAlgoritmo(c *SubAlgoritmoContext)
}
