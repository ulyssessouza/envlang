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
		"", "STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "STR", "SPACE", "CRLF", "ANY",
	}
	staticData.RuleNames = []string{
		"STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "STR", "SPACE", "CRLF", "ANY", "FIRST_CHAR",
		"REST_OF_STRING", "NUMBER", "VAR_ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 125, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 30, 8, 0, 10, 0,
		12, 0, 33, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 41, 8, 0, 10,
		0, 12, 0, 44, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 52, 8, 1,
		10, 1, 12, 1, 55, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 61, 8, 1, 10, 1,
		12, 1, 64, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 72, 8, 2, 10,
		2, 12, 2, 75, 9, 2, 1, 2, 1, 2, 5, 2, 79, 8, 2, 10, 2, 12, 2, 82, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 91, 8, 4, 10, 4, 12, 4,
		94, 9, 4, 1, 5, 1, 5, 1, 6, 3, 6, 99, 8, 6, 1, 6, 1, 6, 3, 6, 103, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 112, 8, 10, 11, 10, 12,
		10, 113, 1, 11, 1, 11, 1, 11, 5, 11, 119, 8, 11, 10, 11, 12, 11, 122, 9,
		11, 3, 11, 124, 8, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 0, 19, 0, 21, 0, 23, 0, 1, 0, 6, 2, 0, 9, 9, 32, 32, 8, 0,
		10, 10, 13, 13, 32, 32, 34, 34, 39, 39, 44, 44, 46, 46, 92, 92, 7, 0, 10,
		10, 13, 13, 32, 32, 34, 34, 36, 36, 39, 39, 92, 92, 1, 0, 48, 57, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 132, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		1, 25, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 85, 1, 0, 0,
		0, 9, 88, 1, 0, 0, 0, 11, 95, 1, 0, 0, 0, 13, 102, 1, 0, 0, 0, 15, 104,
		1, 0, 0, 0, 17, 106, 1, 0, 0, 0, 19, 108, 1, 0, 0, 0, 21, 111, 1, 0, 0,
		0, 23, 123, 1, 0, 0, 0, 25, 26, 5, 36, 0, 0, 26, 27, 5, 123, 0, 0, 27,
		31, 1, 0, 0, 0, 28, 30, 3, 11, 5, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0,
		0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 34, 35, 3, 23, 11, 0, 35, 36, 5, 58, 0, 0, 36, 37, 5, 45, 0,
		0, 37, 38, 1, 0, 0, 0, 38, 42, 3, 9, 4, 0, 39, 41, 3, 11, 5, 0, 40, 39,
		1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 45, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 5, 125, 0, 0, 46, 2, 1,
		0, 0, 0, 47, 48, 5, 36, 0, 0, 48, 49, 5, 123, 0, 0, 49, 53, 1, 0, 0, 0,
		50, 52, 3, 11, 5, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1,
		0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56,
		57, 3, 23, 11, 0, 57, 58, 5, 45, 0, 0, 58, 62, 3, 9, 4, 0, 59, 61, 3, 11,
		5, 0, 60, 59, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63,
		1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 5, 125, 0,
		0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 36, 0, 0, 68, 69, 5, 123, 0, 0, 69, 73,
		1, 0, 0, 0, 70, 72, 3, 11, 5, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0,
		73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1,
		0, 0, 0, 76, 80, 3, 23, 11, 0, 77, 79, 3, 11, 5, 0, 78, 77, 1, 0, 0, 0,
		79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1,
		0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 125, 0, 0, 84, 6, 1, 0, 0, 0, 85,
		86, 5, 36, 0, 0, 86, 87, 3, 23, 11, 0, 87, 8, 1, 0, 0, 0, 88, 92, 3, 17,
		8, 0, 89, 91, 3, 19, 9, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 10, 1, 0, 0, 0, 94, 92, 1, 0, 0,
		0, 95, 96, 7, 0, 0, 0, 96, 12, 1, 0, 0, 0, 97, 99, 5, 13, 0, 0, 98, 97,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 103, 5, 10, 0,
		0, 101, 103, 5, 13, 0, 0, 102, 98, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103,
		14, 1, 0, 0, 0, 104, 105, 9, 0, 0, 0, 105, 16, 1, 0, 0, 0, 106, 107, 8,
		1, 0, 0, 107, 18, 1, 0, 0, 0, 108, 109, 8, 2, 0, 0, 109, 20, 1, 0, 0, 0,
		110, 112, 7, 3, 0, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 22, 1, 0, 0, 0, 115, 124, 3,
		21, 10, 0, 116, 120, 7, 4, 0, 0, 117, 119, 7, 5, 0, 0, 118, 117, 1, 0,
		0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0,
		121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 115, 1, 0, 0, 0, 123,
		116, 1, 0, 0, 0, 124, 24, 1, 0, 0, 0, 13, 0, 31, 42, 53, 62, 73, 80, 92,
		98, 102, 113, 120, 123, 0,
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
	EnvLangValueLexerSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY = 1
	EnvLangValueLexerSTRICT_VAR_WITH_DEFAULT_IF_UNSET          = 2
	EnvLangValueLexerSIMPLE_STRICT_VAR                         = 3
	EnvLangValueLexerSIMPLE_VAR                                = 4
	EnvLangValueLexerSTR                                       = 5
	EnvLangValueLexerSPACE                                     = 6
	EnvLangValueLexerCRLF                                      = 7
	EnvLangValueLexerANY                                       = 8
)
