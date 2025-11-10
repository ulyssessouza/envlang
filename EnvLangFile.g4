grammar EnvLangFile;

// Parser rules
envFile
    : line* EOF
    ;

line
    : entry
    | comment
    | NEWLINE
    ;

entry
    : WS* identifier WS* (ASSIGN WS* value?)? WS* inlineComment? NEWLINE?
    ;

identifier
    : IDENTIFIER
    | UNQUOTED_VALUE
    ;

value
    : DQSTRING
    | SQSTRING
    | IDENTIFIER
    | UNQUOTED_VALUE
    ;

comment
    : COMMENT NEWLINE?
    ;

inlineComment
    : COMMENT
    ;

// Lexer rules

ASSIGN
    : '='
    ;

// Comments: # to end of line
COMMENT
    : '#' ~[\r\n]*
    ;

// Double-quoted string: can span multiple lines, contains everything until closing quote
DQSTRING
    : '"' ( '\\' . | '""' | ~["\r\n] | NEWLINE )* '"'
    ;

// Single-quoted string: can span multiple lines, no variable expansion
SQSTRING
    : '\'' ( '\\' . | '\'\'' | ~['\r\n] | NEWLINE )* '\''
    ;

// Unquoted value: cannot start with whitespace or special chars,
// but can contain spaces and quotes internally. Stops at comment or newline.
// First char: not whitespace, =, #, quotes, or newline
// Following chars: not =, #, or newline (spaces and quotes allowed)
// IMPORTANT: Defined before IDENTIFIER to take precedence
UNQUOTED_VALUE
    : ~[ \t\r\n#="'] ~[\r\n#=]*
    ;

// Identifier: letters, numbers, underscores, dashes, dots
// Must start with letter or underscore
// Defined after UNQUOTED_VALUE, so only matches when UNQUOTED_VALUE doesn't
IDENTIFIER
    : [a-zA-Z_][a-zA-Z0-9_.\-]*
    ;

// Whitespace (spaces and tabs only, not newlines)
WS
    : [ \t]+
    ;

// Newlines
NEWLINE
    : '\r'? '\n'
    | '\r'
    ;
