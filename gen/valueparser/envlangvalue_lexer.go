// Code generated from EnvLangValue.g4 by ANTLR 4.13.1. DO NOT EDIT.

package valueparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EnvLangValueLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EnvLangValueLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func envlangvaluelexerLexerInit() {
	staticData := &EnvLangValueLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'${'", "'}'", "'$'", "' '", "'\"\"'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT_NO_SPACE", "TEXT_ANY",
		"CRLF",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 44, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 3, 8, 39, 8, 8, 1, 8, 1, 8, 3, 8, 43, 8, 8,
		0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0,
		2, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 5, 0, 10, 10, 48, 57, 65, 90,
		95, 95, 97, 122, 45, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 22, 1,
		0, 0, 0, 5, 24, 1, 0, 0, 0, 7, 26, 1, 0, 0, 0, 9, 28, 1, 0, 0, 0, 11, 31,
		1, 0, 0, 0, 13, 33, 1, 0, 0, 0, 15, 35, 1, 0, 0, 0, 17, 42, 1, 0, 0, 0,
		19, 20, 5, 36, 0, 0, 20, 21, 5, 123, 0, 0, 21, 2, 1, 0, 0, 0, 22, 23, 5,
		125, 0, 0, 23, 4, 1, 0, 0, 0, 24, 25, 5, 36, 0, 0, 25, 6, 1, 0, 0, 0, 26,
		27, 5, 32, 0, 0, 27, 8, 1, 0, 0, 0, 28, 29, 5, 34, 0, 0, 29, 30, 5, 34,
		0, 0, 30, 10, 1, 0, 0, 0, 31, 32, 5, 123, 0, 0, 32, 12, 1, 0, 0, 0, 33,
		34, 7, 0, 0, 0, 34, 14, 1, 0, 0, 0, 35, 36, 7, 1, 0, 0, 36, 16, 1, 0, 0,
		0, 37, 39, 5, 13, 0, 0, 38, 37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40,
		1, 0, 0, 0, 40, 43, 5, 10, 0, 0, 41, 43, 5, 13, 0, 0, 42, 38, 1, 0, 0,
		0, 42, 41, 1, 0, 0, 0, 43, 18, 1, 0, 0, 0, 3, 0, 38, 42, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EnvLangValueLexerInit initializes any static state used to implement EnvLangValueLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEnvLangValueLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EnvLangValueLexerInit() {
	staticData := &EnvLangValueLexerLexerStaticData
	staticData.once.Do(envlangvaluelexerLexerInit)
}

// NewEnvLangValueLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEnvLangValueLexer(input antlr.CharStream) *EnvLangValueLexer {
	EnvLangValueLexerInit()
	l := new(EnvLangValueLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EnvLangValueLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
