grammar EnvLangValue;

// Parser rules
dqstring
    : content* EOF
    ;

content
    : variable
    | escapedChar
    | TEXT
    | WS
    | NEWLINE
    ;

variable
    : STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY
    | STRICT_VAR_WITH_DEFAULT_IF_UNSET
    | SIMPLE_STRICT_VAR
    | SIMPLE_VAR
    | DOLLAR  // Lone dollar sign
    ;

escapedChar
    : ESCAPED_CHAR
    ;

// Lexer rules

// Variable substitution with default if unset or empty: ${VAR:-default}
STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY
    : '${' WS* VAR_NAME WS* ':-' DEFAULT_VALUE WS* '}'
    ;

// Variable substitution with default if unset: ${VAR-default}
STRICT_VAR_WITH_DEFAULT_IF_UNSET
    : '${' WS* VAR_NAME WS* '-' DEFAULT_VALUE WS* '}'
    ;

// Simple strict variable: ${VAR}
SIMPLE_STRICT_VAR
    : '${' WS* VAR_NAME WS* '}'
    ;

// Simple variable: $VAR
SIMPLE_VAR
    : '$' VAR_NAME
    ;

// Lone dollar sign (not followed by valid variable name)
DOLLAR
    : '$' ~[a-zA-Z_0-9{]?
    ;

// Escaped characters: \n, \t, \r, \\, \$, \"
ESCAPED_CHAR
    : '\\' [ntr\\$"]
    ;

// Regular text: anything except special chars
TEXT
    : TEXT_CHAR+
    ;

// Whitespace (preserve in values)
WS
    : [ \t]+
    ;

// Newlines (preserve in multiline values)
NEWLINE
    : '\r'? '\n'
    | '\r'
    ;

// Fragment rules

fragment VAR_NAME
    : [a-zA-Z_][a-zA-Z_0-9]*  // Standard variable names
    | [0-9]+                  // Numeric variables ($1, $2, etc.)
    ;

fragment DEFAULT_VALUE
    : ( ~[}] )*  // Everything until closing brace
    ;

fragment TEXT_CHAR
    : ~[$\\\r\n \t]  // Anything except dollar, backslash, newline, or whitespace
    ;
