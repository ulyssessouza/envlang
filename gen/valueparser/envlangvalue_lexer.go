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
	staticData.SymbolicNames = []string{
		"", "STRICT_VAR", "SIMPLE_VAR", "STR", "SPACE", "CRLF", "ANY",
	}
	staticData.RuleNames = []string{
		"STRICT_VAR", "SIMPLE_VAR", "STR", "SPACE", "CRLF", "FIRST_CHAR", "REST_OF_STRING",
		"VAR_ID", "ANY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 76, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		0, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 0, 1, 0, 5, 0, 31,
		8, 0, 10, 0, 12, 0, 34, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		5, 2, 43, 8, 2, 10, 2, 12, 2, 46, 9, 2, 1, 3, 1, 3, 1, 4, 3, 4, 51, 8,
		4, 1, 4, 1, 4, 3, 4, 55, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 4, 7, 62,
		8, 7, 11, 7, 12, 7, 63, 1, 7, 1, 7, 5, 7, 68, 8, 7, 10, 7, 12, 7, 71, 9,
		7, 3, 7, 73, 8, 7, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		0, 13, 0, 15, 0, 17, 6, 1, 0, 6, 2, 0, 9, 9, 32, 32, 8, 0, 10, 10, 13,
		13, 32, 32, 34, 34, 39, 39, 44, 44, 46, 46, 92, 92, 7, 0, 10, 10, 13, 13,
		32, 32, 34, 34, 36, 36, 39, 39, 92, 92, 1, 0, 48, 57, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 80, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 40, 1, 0,
		0, 0, 7, 47, 1, 0, 0, 0, 9, 54, 1, 0, 0, 0, 11, 56, 1, 0, 0, 0, 13, 58,
		1, 0, 0, 0, 15, 72, 1, 0, 0, 0, 17, 74, 1, 0, 0, 0, 19, 20, 5, 36, 0, 0,
		20, 21, 5, 123, 0, 0, 21, 25, 1, 0, 0, 0, 22, 24, 3, 7, 3, 0, 23, 22, 1,
		0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26,
		28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 32, 3, 15, 7, 0, 29, 31, 3, 7,
		3, 0, 30, 29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33,
		1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 5, 125, 0,
		0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 36, 0, 0, 38, 39, 3, 15, 7, 0, 39, 4,
		1, 0, 0, 0, 40, 44, 3, 11, 5, 0, 41, 43, 3, 13, 6, 0, 42, 41, 1, 0, 0,
		0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 6, 1,
		0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 7, 0, 0, 0, 48, 8, 1, 0, 0, 0, 49,
		51, 5, 13, 0, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0,
		0, 0, 52, 55, 5, 10, 0, 0, 53, 55, 5, 13, 0, 0, 54, 50, 1, 0, 0, 0, 54,
		53, 1, 0, 0, 0, 55, 10, 1, 0, 0, 0, 56, 57, 8, 1, 0, 0, 57, 12, 1, 0, 0,
		0, 58, 59, 8, 2, 0, 0, 59, 14, 1, 0, 0, 0, 60, 62, 7, 3, 0, 0, 61, 60,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0,
		64, 73, 1, 0, 0, 0, 65, 69, 7, 4, 0, 0, 66, 68, 7, 5, 0, 0, 67, 66, 1,
		0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70,
		73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 61, 1, 0, 0, 0, 72, 65, 1, 0, 0,
		0, 73, 16, 1, 0, 0, 0, 74, 75, 9, 0, 0, 0, 75, 18, 1, 0, 0, 0, 9, 0, 25,
		32, 44, 50, 54, 63, 69, 72, 0,
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
	EnvLangValueLexerSTRICT_VAR = 1
	EnvLangValueLexerSIMPLE_VAR = 2
	EnvLangValueLexerSTR        = 3
	EnvLangValueLexerSPACE      = 4
	EnvLangValueLexerCRLF       = 5
	EnvLangValueLexerANY        = 6
)
