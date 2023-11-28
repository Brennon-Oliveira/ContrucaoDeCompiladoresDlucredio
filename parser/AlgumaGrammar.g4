lexer grammar AlgumaGrammar;

PalavraChave: 'DECLARACOES' | 'ALGORITMO' | 'INTEIRO' | 'REAL' | 'ATRIBUIR' | 'A' | 'LER' |
                'IMPRIMIR' | 'SE' | 'ENTAO' | 'ENQUANTO' | 'INICIO' | 'FIM' | 'E' | 'OU';

NumInt: ('+' | '-')? (Digito)+;

NumReal: ('+' | '-')? (Digito)+ '.' (Digito)+;

Variavel: Letra (Letra | Digito)*;

Cadeia : '\'' (ESC_SEQ | ~('\\'))* '\'';

fragment
ESC_SEQ : '\\\'';

OP_REL: '>' | '<' | '>=' | '<=' | '==' | '<>';

OP_ARIT: '+' | '-' | '*' | '/';

DELIM: ':';

ABREPAR: '(';

FECHAPAR: ')';

WS: ( ' ' | '\t' | '\r' | '\n' ) -> skip;

Letra: 'a'..'z' | 'A'..'Z';
Digito: '0'..'9';