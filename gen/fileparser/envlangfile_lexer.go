// Code generated from EnvLangFile.g4 by ANTLR 4.13.1. DO NOT EDIT.

package fileparser

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

type EnvLangFileLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EnvLangFileLexerLexerStaticData struct {
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

func envlangfilelexerLexerInit() {
	staticData := &EnvLangFileLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'\\r'", "'\\n'", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 54, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 5, 3, 24, 8, 3, 10, 3, 12, 3, 27, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 33, 8, 4, 10, 4, 12, 4, 36, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 44, 8, 5, 10, 5, 12, 5, 47, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 0, 0, 7, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 4, 7,
		0, 10, 10, 13, 13, 32, 32, 34, 34, 39, 39, 44, 44, 61, 61, 6, 0, 10, 10,
		13, 13, 34, 34, 39, 39, 44, 44, 61, 61, 1, 0, 34, 34, 1, 0, 39, 39, 58,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0,
		0, 3, 17, 1, 0, 0, 0, 5, 19, 1, 0, 0, 0, 7, 21, 1, 0, 0, 0, 9, 28, 1, 0,
		0, 0, 11, 39, 1, 0, 0, 0, 13, 50, 1, 0, 0, 0, 15, 16, 5, 61, 0, 0, 16,
		2, 1, 0, 0, 0, 17, 18, 5, 13, 0, 0, 18, 4, 1, 0, 0, 0, 19, 20, 5, 10, 0,
		0, 20, 6, 1, 0, 0, 0, 21, 25, 8, 0, 0, 0, 22, 24, 8, 1, 0, 0, 23, 22, 1,
		0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26,
		8, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 34, 5, 34, 0, 0, 29, 30, 5, 34,
		0, 0, 30, 33, 5, 34, 0, 0, 31, 33, 8, 2, 0, 0, 32, 29, 1, 0, 0, 0, 32,
		31, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0,
		0, 35, 37, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 38, 5, 34, 0, 0, 38, 10,
		1, 0, 0, 0, 39, 45, 5, 39, 0, 0, 40, 41, 5, 39, 0, 0, 41, 44, 5, 39, 0,
		0, 42, 44, 8, 3, 0, 0, 43, 40, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 47,
		1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 48, 49, 5, 39, 0, 0, 49, 12, 1, 0, 0, 0, 50, 51, 5,
		32, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 6, 0, 0, 53, 14, 1, 0, 0, 0, 6,
		0, 25, 32, 34, 43, 45, 1, 6, 0, 0,
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

// EnvLangFileLexerInit initializes any static state used to implement EnvLangFileLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEnvLangFileLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EnvLangFileLexerInit() {
	staticData := &EnvLangFileLexerLexerStaticData
	staticData.once.Do(envlangfilelexerLexerInit)
}

// NewEnvLangFileLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEnvLangFileLexer(input antlr.CharStream) *EnvLangFileLexer {
	EnvLangFileLexerInit()
	l := new(EnvLangFileLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EnvLangFileLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
