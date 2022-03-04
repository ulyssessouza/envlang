// Code generated from EnvLangValue.g4 by ANTLR 4.9.3. DO NOT EDIT.

package valueparser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 46, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 5, 10, 41, 10, 10, 3, 10, 3, 10, 5, 10, 45, 10, 10, 2, 2,
	11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2,
	4, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 7, 2, 12, 12, 50, 59, 67, 92,
	97, 97, 99, 124, 2, 47, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 24,
	3, 2, 2, 2, 7, 26, 3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 30, 3, 2, 2, 2, 13,
	33, 3, 2, 2, 2, 15, 35, 3, 2, 2, 2, 17, 37, 3, 2, 2, 2, 19, 44, 3, 2, 2,
	2, 21, 22, 7, 38, 2, 2, 22, 23, 7, 125, 2, 2, 23, 4, 3, 2, 2, 2, 24, 25,
	7, 127, 2, 2, 25, 6, 3, 2, 2, 2, 26, 27, 7, 38, 2, 2, 27, 8, 3, 2, 2, 2,
	28, 29, 7, 34, 2, 2, 29, 10, 3, 2, 2, 2, 30, 31, 7, 36, 2, 2, 31, 32, 7,
	36, 2, 2, 32, 12, 3, 2, 2, 2, 33, 34, 7, 125, 2, 2, 34, 14, 3, 2, 2, 2,
	35, 36, 9, 2, 2, 2, 36, 16, 3, 2, 2, 2, 37, 38, 9, 3, 2, 2, 38, 18, 3,
	2, 2, 2, 39, 41, 7, 15, 2, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41,
	42, 3, 2, 2, 2, 42, 45, 7, 12, 2, 2, 43, 45, 7, 15, 2, 2, 44, 40, 3, 2,
	2, 2, 44, 43, 3, 2, 2, 2, 45, 20, 3, 2, 2, 2, 5, 2, 40, 44, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'${'", "'}'", "'$'", "' '", "'\"\"'", "'{'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT_NO_SPACE", "TEXT_ANY",
	"CRLF",
}

type EnvLangValueLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewEnvLangValueLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *EnvLangValueLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEnvLangValueLexer(input antlr.CharStream) *EnvLangValueLexer {
	l := new(EnvLangValueLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EnvLangValue.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EnvLangValueLexer tokens.
const (
	EnvLangValueLexerT__0          = 1
	EnvLangValueLexerT__1          = 2
	EnvLangValueLexerT__2          = 3
	EnvLangValueLexerT__3          = 4
	EnvLangValueLexerT__4          = 5
	EnvLangValueLexerT__5          = 6
	EnvLangValueLexerTEXT_NO_SPACE = 7
	EnvLangValueLexerTEXT_ANY      = 8
	EnvLangValueLexerCRLF          = 9
)
