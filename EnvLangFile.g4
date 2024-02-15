grammar EnvLangFile;

envFile: NL* (entry NL*)+ ;

entry : identifier (ASSIGN value | ASSIGN)? (NL* | EOF) ;

identifier : TEXT;

value : op=(DQSTRING | SQSTRING | TEXT) ;

ASSIGN : '=';
NL : '\r'? '\n';
TEXT   : ~[=,\r\n"' ]~[=,\r\n"']* ;
DQSTRING : '"' ('""'|~'"')* '"' ;
SQSTRING : '\'' ('\'\''|~'\'')* '\'' ;

SPACE : ' ' -> skip;
