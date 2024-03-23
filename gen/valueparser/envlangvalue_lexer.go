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
		"", "", "", "", "", "", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "STR", "PESO_SIGN", "FIRST_CHAR",
		"REST_OF_STRING", "NUMBER", "VAR_ID", "SPACE", "CRLF", "ANY",
	}
	staticData.RuleNames = []string{
		"STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "STR", "PESO_SIGN", "FIRST_CHAR",
		"REST_OF_STRING", "NUMBER", "VAR_ID", "SPACE", "CRLF", "ANY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 129, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 32,
		8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 63,
		8, 1, 10, 1, 12, 1, 66, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2,
		74, 8, 2, 10, 2, 12, 2, 77, 9, 2, 1, 2, 1, 2, 5, 2, 81, 8, 2, 10, 2, 12,
		2, 84, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 93, 8, 4,
		10, 4, 12, 4, 96, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8,
		105, 8, 8, 11, 8, 12, 8, 106, 1, 9, 1, 9, 1, 9, 5, 9, 112, 8, 9, 10, 9,
		12, 9, 115, 9, 9, 3, 9, 117, 8, 9, 1, 10, 1, 10, 1, 11, 3, 11, 122, 8,
		11, 1, 11, 1, 11, 3, 11, 126, 8, 11, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 1, 0, 6, 8, 0, 10, 10, 13, 13, 32, 32, 34, 34, 39, 39, 44, 44, 46,
		46, 92, 92, 7, 0, 10, 10, 13, 13, 32, 32, 34, 34, 36, 36, 39, 39, 92, 92,
		1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 2, 0, 9, 9, 32, 32, 140, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1,
		0, 0, 0, 3, 49, 1, 0, 0, 0, 5, 69, 1, 0, 0, 0, 7, 87, 1, 0, 0, 0, 9, 90,
		1, 0, 0, 0, 11, 97, 1, 0, 0, 0, 13, 99, 1, 0, 0, 0, 15, 101, 1, 0, 0, 0,
		17, 104, 1, 0, 0, 0, 19, 116, 1, 0, 0, 0, 21, 118, 1, 0, 0, 0, 23, 125,
		1, 0, 0, 0, 25, 127, 1, 0, 0, 0, 27, 28, 5, 36, 0, 0, 28, 29, 5, 123, 0,
		0, 29, 33, 1, 0, 0, 0, 30, 32, 3, 21, 10, 0, 31, 30, 1, 0, 0, 0, 32, 35,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 36, 37, 3, 19, 9, 0, 37, 38, 5, 58, 0, 0, 38, 39, 5,
		45, 0, 0, 39, 40, 1, 0, 0, 0, 40, 44, 3, 9, 4, 0, 41, 43, 3, 21, 10, 0,
		42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1,
		0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 125, 0, 0,
		48, 2, 1, 0, 0, 0, 49, 50, 5, 36, 0, 0, 50, 51, 5, 123, 0, 0, 51, 55, 1,
		0, 0, 0, 52, 54, 3, 21, 10, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 55, 1,
		0, 0, 0, 58, 59, 3, 19, 9, 0, 59, 60, 5, 45, 0, 0, 60, 64, 3, 9, 4, 0,
		61, 63, 3, 21, 10, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67,
		68, 5, 125, 0, 0, 68, 4, 1, 0, 0, 0, 69, 70, 5, 36, 0, 0, 70, 71, 5, 123,
		0, 0, 71, 75, 1, 0, 0, 0, 72, 74, 3, 21, 10, 0, 73, 72, 1, 0, 0, 0, 74,
		77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0,
		0, 77, 75, 1, 0, 0, 0, 78, 82, 3, 19, 9, 0, 79, 81, 3, 21, 10, 0, 80, 79,
		1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 125, 0, 0, 86, 6, 1,
		0, 0, 0, 87, 88, 5, 36, 0, 0, 88, 89, 3, 19, 9, 0, 89, 8, 1, 0, 0, 0, 90,
		94, 3, 13, 6, 0, 91, 93, 3, 15, 7, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0,
		0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 10, 1, 0, 0, 0, 96, 94,
		1, 0, 0, 0, 97, 98, 5, 36, 0, 0, 98, 12, 1, 0, 0, 0, 99, 100, 8, 0, 0,
		0, 100, 14, 1, 0, 0, 0, 101, 102, 8, 1, 0, 0, 102, 16, 1, 0, 0, 0, 103,
		105, 7, 2, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 18, 1, 0, 0, 0, 108, 117, 3, 17,
		8, 0, 109, 113, 7, 3, 0, 0, 110, 112, 7, 4, 0, 0, 111, 110, 1, 0, 0, 0,
		112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 108, 1, 0, 0, 0, 116, 109,
		1, 0, 0, 0, 117, 20, 1, 0, 0, 0, 118, 119, 7, 5, 0, 0, 119, 22, 1, 0, 0,
		0, 120, 122, 5, 13, 0, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 126, 5, 10, 0, 0, 124, 126, 5, 13, 0, 0, 125, 121,
		1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126, 24, 1, 0, 0, 0, 127, 128, 9, 0,
		0, 0, 128, 26, 1, 0, 0, 0, 13, 0, 33, 44, 55, 64, 75, 82, 94, 106, 113,
		116, 121, 125, 0,
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
	EnvLangValueLexerPESO_SIGN                                 = 6
	EnvLangValueLexerFIRST_CHAR                                = 7
	EnvLangValueLexerREST_OF_STRING                            = 8
	EnvLangValueLexerNUMBER                                    = 9
	EnvLangValueLexerVAR_ID                                    = 10
	EnvLangValueLexerSPACE                                     = 11
	EnvLangValueLexerCRLF                                      = 12
	EnvLangValueLexerANY                                       = 13
)
