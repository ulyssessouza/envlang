grammar EnvLangFile;

envFile: nl* (entry nl*)+ ;

entry : identifier (assign value | assign)? (nl* | EOF) ;

//TEXT_NO_SPACE : [a-zA-Z0-9_] ;
identifier : TEXT;

assign : '=';

value : op=(DQSTRING | SQSTRING | TEXT) ;

nl : '\r'? '\n';

TEXT   : ~[=,\r\n"' ]~[=,\r\n"']* ;
DQSTRING : '"' ('""'|~'"')* '"' ;
SQSTRING : '\'' ('\'\''|~'\'')* '\'' ;

SPACE : ' ' -> skip;