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
		"", "ASSIGN", "CRLF", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"ASSIGN", "CRLF", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 55, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 3, 1, 17, 8, 1, 1, 1, 1, 1, 3, 1,
		21, 8, 1, 1, 2, 1, 2, 5, 2, 25, 8, 2, 10, 2, 12, 2, 28, 9, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 34, 8, 3, 10, 3, 12, 3, 37, 9, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 45, 8, 4, 10, 4, 12, 4, 48, 9, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 0, 0, 6, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0,
		4, 7, 0, 10, 10, 13, 13, 32, 32, 34, 34, 39, 39, 44, 44, 61, 61, 6, 0,
		10, 10, 13, 13, 34, 34, 39, 39, 44, 44, 61, 61, 1, 0, 34, 34, 1, 0, 39,
		39, 61, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 1, 13, 1, 0, 0, 0, 3, 20,
		1, 0, 0, 0, 5, 22, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0, 9, 40, 1, 0, 0, 0, 11,
		51, 1, 0, 0, 0, 13, 14, 5, 61, 0, 0, 14, 2, 1, 0, 0, 0, 15, 17, 5, 13,
		0, 0, 16, 15, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 21,
		5, 10, 0, 0, 19, 21, 5, 13, 0, 0, 20, 16, 1, 0, 0, 0, 20, 19, 1, 0, 0,
		0, 21, 4, 1, 0, 0, 0, 22, 26, 8, 0, 0, 0, 23, 25, 8, 1, 0, 0, 24, 23, 1,
		0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27,
		6, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 29, 35, 5, 34, 0, 0, 30, 31, 5, 34,
		0, 0, 31, 34, 5, 34, 0, 0, 32, 34, 8, 2, 0, 0, 33, 30, 1, 0, 0, 0, 33,
		32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0,
		0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 34, 0, 0, 39, 8,
		1, 0, 0, 0, 40, 46, 5, 39, 0, 0, 41, 42, 5, 39, 0, 0, 42, 45, 5, 39, 0,
		0, 43, 45, 8, 3, 0, 0, 44, 41, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 48,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 49, 50, 5, 39, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52, 5,
		32, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 6, 5, 0, 0, 54, 12, 1, 0, 0, 0, 8,
		0, 16, 20, 26, 33, 35, 44, 46, 1, 6, 0, 0,
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
	EnvLangFileLexerCRLF     = 2
	EnvLangFileLexerTEXT     = 3
	EnvLangFileLexerDQSTRING = 4
	EnvLangFileLexerSQSTRING = 5
	EnvLangFileLexerSPACE    = 6
)
