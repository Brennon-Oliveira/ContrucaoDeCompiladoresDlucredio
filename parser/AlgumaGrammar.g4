lexer grammar AlgumaGrammar;

Letra: 'a'..'z' | 'A'..'Z';
Digito: '0'..'9';
Variavel: Letra (Letra | Digito)*;