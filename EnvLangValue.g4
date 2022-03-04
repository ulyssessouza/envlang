grammar EnvLangValue;

dqstring : content* (CRLF* | EOF)  ;

strictVar : '${' TEXT_NO_SPACE* '}';
simpleVar : '$' TEXT_NO_SPACE* ;

variable : strictVar | simpleVar ;

TEXT_NO_SPACE : [a-zA-Z0-9_] ;
TEXT_ANY      : [a-zA-Z0-9_\n] ;

text : TEXT_NO_SPACE TEXT_ANY* ;

space : ' ';

dQEscape : '""' ;

special
    : '{'
    | '}'
    ;

content
    : dQEscape
    | variable
    | space
    | special
    | text
    ;

CRLF
    : '\r'? '\n'
    | '\r'
    ;
