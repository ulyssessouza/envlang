grammar EnvLangValue;

dqstring
	: content* (SPACE* | CRLF* | EOF)
	;

content
	: variable
	| STR
	| SPACE
	| CRLF
	;

variable
	: var=
	(
	  STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY
	| STRICT_VAR_WITH_DEFAULT_IF_UNSET
	| SIMPLE_STRICT_VAR
	| SIMPLE_VAR
	)
	;

STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY
	: '${' SPACE* VAR_ID ':-' STR SPACE* '}'
	;

STRICT_VAR_WITH_DEFAULT_IF_UNSET
	: '${' SPACE* VAR_ID '-' STR SPACE* '}'
	;

SIMPLE_STRICT_VAR
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

// Catch all rule.
// This is used to avoid the error message "no viable alternative at input '<EOF>'". Just for debugging.
ANY
	: .
	;

fragment FIRST_CHAR
	: ~[,\\."'\r\n ]
	;

fragment REST_OF_STRING
	: ~[\\'"$\r\n ]
	;

fragment NUMBER
	: [0-9]+
	;

fragment VAR_ID
	: NUMBER // Only numbers
	| [a-zA-Z_][a-zA-Z_0-9]* // Start with letters, then letters, numbers and underscores
	;
