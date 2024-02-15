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
		"", "'='", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "NL", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"ASSIGN", "NL", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 53, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 3, 1, 17, 8, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 5, 2, 23, 8, 2, 10, 2, 12, 2, 26, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 32, 8, 3, 10, 3, 12, 3, 35, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 43, 8, 4, 10, 4, 12, 4, 46, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 0, 0, 6, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 4, 7, 0, 10, 10,
		13, 13, 32, 32, 34, 34, 39, 39, 44, 44, 61, 61, 6, 0, 10, 10, 13, 13, 34,
		34, 39, 39, 44, 44, 61, 61, 1, 0, 34, 34, 1, 0, 39, 39, 58, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 1, 13, 1, 0, 0, 0, 3, 16, 1, 0, 0, 0, 5, 20, 1,
		0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13,
		14, 5, 61, 0, 0, 14, 2, 1, 0, 0, 0, 15, 17, 5, 13, 0, 0, 16, 15, 1, 0,
		0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 19, 5, 10, 0, 0, 19,
		4, 1, 0, 0, 0, 20, 24, 8, 0, 0, 0, 21, 23, 8, 1, 0, 0, 22, 21, 1, 0, 0,
		0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 6, 1,
		0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 33, 5, 34, 0, 0, 28, 29, 5, 34, 0, 0,
		29, 32, 5, 34, 0, 0, 30, 32, 8, 2, 0, 0, 31, 28, 1, 0, 0, 0, 31, 30, 1,
		0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34,
		36, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 37, 5, 34, 0, 0, 37, 8, 1, 0, 0,
		0, 38, 44, 5, 39, 0, 0, 39, 40, 5, 39, 0, 0, 40, 43, 5, 39, 0, 0, 41, 43,
		8, 3, 0, 0, 42, 39, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0,
		44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1,
		0, 0, 0, 47, 48, 5, 39, 0, 0, 48, 10, 1, 0, 0, 0, 49, 50, 5, 32, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 52, 6, 5, 0, 0, 52, 12, 1, 0, 0, 0, 7, 0, 16, 24,
		31, 33, 42, 44, 1, 6, 0, 0,
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
	EnvLangFileLexerASSIGN   = 1
	EnvLangFileLexerNL       = 2
	EnvLangFileLexerTEXT     = 3
	EnvLangFileLexerDQSTRING = 4
	EnvLangFileLexerSQSTRING = 5
	EnvLangFileLexerSPACE    = 6
)
