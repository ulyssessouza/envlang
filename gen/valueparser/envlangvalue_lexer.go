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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 42, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 5, 9, 37, 10,
	9, 3, 9, 3, 9, 5, 9, 41, 10, 9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 3, 2, 4, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 7,
	2, 12, 12, 50, 59, 67, 92, 97, 97, 99, 124, 2, 43, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2,
	5, 22, 3, 2, 2, 2, 7, 24, 3, 2, 2, 2, 9, 26, 3, 2, 2, 2, 11, 28, 3, 2,
	2, 2, 13, 31, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 40, 3, 2, 2, 2, 19, 20,
	7, 38, 2, 2, 20, 21, 7, 125, 2, 2, 21, 4, 3, 2, 2, 2, 22, 23, 7, 127, 2,
	2, 23, 6, 3, 2, 2, 2, 24, 25, 7, 38, 2, 2, 25, 8, 3, 2, 2, 2, 26, 27, 7,
	34, 2, 2, 27, 10, 3, 2, 2, 2, 28, 29, 7, 36, 2, 2, 29, 30, 7, 36, 2, 2,
	30, 12, 3, 2, 2, 2, 31, 32, 9, 2, 2, 2, 32, 14, 3, 2, 2, 2, 33, 34, 9,
	3, 2, 2, 34, 16, 3, 2, 2, 2, 35, 37, 7, 15, 2, 2, 36, 35, 3, 2, 2, 2, 36,
	37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 41, 7, 12, 2, 2, 39, 41, 7, 15,
	2, 2, 40, 36, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 18, 3, 2, 2, 2, 5, 2,
	36, 40, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'${'", "'}'", "'$'", "' '", "'\"\"'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
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
	EnvLangValueLexerTEXT_NO_SPACE = 6
	EnvLangValueLexerTEXT_ANY      = 7
	EnvLangValueLexerCRLF          = 8
)
