grammar EnvLangFile;

envFile
	: CRLF* (entry CRLF*)+
	;

entry
	: identifier (ASSIGN value | ASSIGN)? (SPACE* CRLF* | EOF)
	;

identifier
	: TEXT
	;

value
	: str=(DQSTRING | SQSTRING | TEXT)
	;

ASSIGN
	: '='
	;

CRLF
	: '\r'? '\n'
	| '\r'
	;

TEXT
	: ~[=,\r\n"' ]~[=,\r\n"']*
	;

DQSTRING
	: '"' ('""' | ~'"')* '"'
	;

SQSTRING
	: '\'' ('\'\'' | ~'\'')* '\''
	;

SPACE
	: ' ' -> skip
	;
