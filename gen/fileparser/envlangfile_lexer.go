// Code generated from EnvLangFile.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "COMMENT", "DQSTRING", "SQSTRING", "UNQUOTED_VALUE", "IDENTIFIER",
		"WS", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"ASSIGN", "COMMENT", "DQSTRING", "SQSTRING", "UNQUOTED_VALUE", "IDENTIFIER",
		"WS", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 80, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 5,
		1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 34, 8, 2, 10, 2, 12, 2, 37, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 48, 8, 3, 10, 3, 12, 3, 51, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 5, 4, 57, 8, 4, 10, 4, 12, 4, 60, 9, 4, 1, 5, 1, 5, 5,
		5, 64, 8, 5, 10, 5, 12, 5, 67, 9, 5, 1, 6, 4, 6, 70, 8, 6, 11, 6, 12, 6,
		71, 1, 7, 3, 7, 75, 8, 7, 1, 7, 1, 7, 3, 7, 79, 8, 7, 0, 0, 8, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 8, 2, 0, 10, 10, 13, 13,
		3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 10, 10, 13, 13, 39, 39, 6, 0, 9, 10,
		13, 13, 32, 32, 34, 35, 39, 39, 61, 61, 4, 0, 10, 10, 13, 13, 35, 35, 61,
		61, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 45, 46, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 9, 9, 32, 32, 93, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 19, 1, 0, 0, 0,
		5, 26, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9, 54, 1, 0, 0, 0, 11, 61, 1, 0,
		0, 0, 13, 69, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17, 18, 5, 61, 0, 0, 18,
		2, 1, 0, 0, 0, 19, 23, 5, 35, 0, 0, 20, 22, 8, 0, 0, 0, 21, 20, 1, 0, 0,
		0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 4, 1,
		0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 35, 5, 34, 0, 0, 27, 28, 5, 92, 0, 0,
		28, 34, 9, 0, 0, 0, 29, 30, 5, 34, 0, 0, 30, 34, 5, 34, 0, 0, 31, 34, 8,
		1, 0, 0, 32, 34, 3, 15, 7, 0, 33, 27, 1, 0, 0, 0, 33, 29, 1, 0, 0, 0, 33,
		31, 1, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0,
		0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39,
		5, 34, 0, 0, 39, 6, 1, 0, 0, 0, 40, 49, 5, 39, 0, 0, 41, 42, 5, 92, 0,
		0, 42, 48, 9, 0, 0, 0, 43, 44, 5, 39, 0, 0, 44, 48, 5, 39, 0, 0, 45, 48,
		8, 2, 0, 0, 46, 48, 3, 15, 7, 0, 47, 41, 1, 0, 0, 0, 47, 43, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1,
		0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52,
		53, 5, 39, 0, 0, 53, 8, 1, 0, 0, 0, 54, 58, 8, 3, 0, 0, 55, 57, 8, 4, 0,
		0, 56, 55, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 10, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 65, 7, 5, 0, 0,
		62, 64, 7, 6, 0, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 12, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68,
		70, 7, 7, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0,
		0, 71, 72, 1, 0, 0, 0, 72, 14, 1, 0, 0, 0, 73, 75, 5, 13, 0, 0, 74, 73,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 79, 5, 10, 0, 0,
		77, 79, 5, 13, 0, 0, 78, 74, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 16, 1,
		0, 0, 0, 11, 0, 23, 33, 35, 47, 49, 58, 65, 71, 74, 78, 0,
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
	EnvLangFileLexerASSIGN         = 1
	EnvLangFileLexerCOMMENT        = 2
	EnvLangFileLexerDQSTRING       = 3
	EnvLangFileLexerSQSTRING       = 4
	EnvLangFileLexerUNQUOTED_VALUE = 5
	EnvLangFileLexerIDENTIFIER     = 6
	EnvLangFileLexerWS             = 7
	EnvLangFileLexerNEWLINE        = 8
)
