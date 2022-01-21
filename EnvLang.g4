grammar EnvLang;

envFile: nl* (entry nl*)+ ;

entry : identifier (assign value | assign)? (nl* | EOF) ;

identifier : TEXT;

assign : '=';

value : op=(DQSTRING | SQSTRING | TEXT) ;

nl : '\r'? '\n';

TEXT   : ~[=,\r\n"' ]~[=,\r\n"']* ;
DQSTRING : '"' ('""'|~'"')* '"' ;
SQSTRING : '\'' ('\'\''|~'\'')* '\'' ;

SPACE : ' ' -> skip;