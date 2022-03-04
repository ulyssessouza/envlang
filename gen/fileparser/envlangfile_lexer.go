// Code generated from EnvLangFile.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fileparser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 56, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 26,
	10, 5, 12, 5, 14, 5, 29, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 35, 10, 6,
	12, 6, 14, 6, 38, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 46,
	10, 7, 12, 7, 14, 7, 49, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 2,
	2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 6, 9, 2, 12, 12,
	15, 15, 34, 34, 36, 36, 41, 41, 46, 46, 63, 63, 8, 2, 12, 12, 15, 15, 36,
	36, 41, 41, 46, 46, 63, 63, 3, 2, 36, 36, 3, 2, 41, 41, 2, 60, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19,
	3, 2, 2, 2, 7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 30, 3, 2, 2, 2, 13,
	41, 3, 2, 2, 2, 15, 52, 3, 2, 2, 2, 17, 18, 7, 63, 2, 2, 18, 4, 3, 2, 2,
	2, 19, 20, 7, 15, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 7, 12, 2, 2, 22, 8,
	3, 2, 2, 2, 23, 27, 10, 2, 2, 2, 24, 26, 10, 3, 2, 2, 25, 24, 3, 2, 2,
	2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 10,
	3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 36, 7, 36, 2, 2, 31, 32, 7, 36, 2,
	2, 32, 35, 7, 36, 2, 2, 33, 35, 10, 4, 2, 2, 34, 31, 3, 2, 2, 2, 34, 33,
	3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2,
	37, 39, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 40, 7, 36, 2, 2, 40, 12, 3,
	2, 2, 2, 41, 47, 7, 41, 2, 2, 42, 43, 7, 41, 2, 2, 43, 46, 7, 41, 2, 2,
	44, 46, 10, 5, 2, 2, 45, 42, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3,
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49,
	47, 3, 2, 2, 2, 50, 51, 7, 41, 2, 2, 51, 14, 3, 2, 2, 2, 52, 53, 7, 34,
	2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 8, 8, 2, 2, 55, 16, 3, 2, 2, 2, 8, 2,
	27, 34, 36, 45, 47, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'\r'", "'\n'", "", "", "", "' '",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
}

type EnvLangFileLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewEnvLangFileLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *EnvLangFileLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEnvLangFileLexer(input antlr.CharStream) *EnvLangFileLexer {
	l := new(EnvLangFileLexer)
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
	l.GrammarFileName = "EnvLangFile.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EnvLangFileLexer tokens.
const (
	EnvLangFileLexerT__0     = 1
	EnvLangFileLexerT__1     = 2
	EnvLangFileLexerT__2     = 3
	EnvLangFileLexerTEXT     = 4
	EnvLangFileLexerDQSTRING = 5
	EnvLangFileLexerSQSTRING = 6
	EnvLangFileLexerSPACE    = 7
)
