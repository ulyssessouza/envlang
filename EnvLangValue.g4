grammar EnvLangValue;

dqstring
	: content* (SPACE* | CRLF* | EOF)
	;

variable
	: var=(STRICT_VAR | SIMPLE_VAR)
	;

content
	: variable
	| STR
	| SPACE
	| CRLF
	;

STRICT_VAR
	: '${' SPACE* VAR_ID SPACE* '}'
	;

SIMPLE_VAR
	: '$' VAR_ID
	;

STR
	: FIRST_CHAR REST_OF_STRING*
	;

SPACE
	: ' '
	| '\t'
	;

CRLF
	: '\r'? '\n'
	| '\r'
	;

fragment FIRST_CHAR
	: ~[,\\."'\r\n ]
	;

fragment REST_OF_STRING
	: ~[\\'"$\r\n ]
	;

fragment VAR_ID
	: [0-9]+ // Only numbers
	| [a-zA-Z_][a-zA-Z_0-9]* // Start with letters, then letters, numbers and underscores
	;

// Catch all rule.
// This is used to avoid the error message "no viable alternative at input '<EOF>'". Just for debugging.
ANY
	: .
	;

