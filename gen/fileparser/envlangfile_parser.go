// Code generated from EnvLangFile.g4 by ANTLR 4.13.1. DO NOT EDIT.

package fileparser // EnvLangFile
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EnvLangFileParser struct {
	*antlr.BaseParser
}

var EnvLangFileParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func envlangfileParserInit() {
	staticData := &EnvLangFileParserStaticData
	staticData.LiteralNames = []string{
		"", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "COMMENT", "DQSTRING", "SQSTRING", "UNQUOTED_VALUE", "IDENTIFIER",
		"WS", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"envFile", "line", "entry", "identifier", "value", "comment", "inlineComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 75, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 2, 5, 2, 29, 8, 2, 10, 2,
		12, 2, 32, 9, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1,
		2, 1, 2, 5, 2, 43, 8, 2, 10, 2, 12, 2, 46, 9, 2, 1, 2, 3, 2, 49, 8, 2,
		3, 2, 51, 8, 2, 1, 2, 5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 3,
		2, 60, 8, 2, 1, 2, 3, 2, 63, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		3, 5, 71, 8, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 2,
		1, 0, 5, 6, 1, 0, 3, 6, 79, 0, 17, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 30,
		1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12,
		72, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0, 0, 16, 19, 1, 0, 0,
		0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 20, 1, 0, 0, 0, 19, 17,
		1, 0, 0, 0, 20, 21, 5, 0, 0, 1, 21, 1, 1, 0, 0, 0, 22, 26, 3, 4, 2, 0,
		23, 26, 3, 10, 5, 0, 24, 26, 5, 8, 0, 0, 25, 22, 1, 0, 0, 0, 25, 23, 1,
		0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 3, 1, 0, 0, 0, 27, 29, 5, 7, 0, 0, 28,
		27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0,
		0, 31, 33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 37, 3, 6, 3, 0, 34, 36,
		5, 7, 0, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 50, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 44, 5,
		1, 0, 0, 41, 43, 5, 7, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44,
		42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0,
		0, 47, 49, 3, 8, 4, 0, 48, 47, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51,
		1, 0, 0, 0, 50, 40, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 55, 1, 0, 0, 0,
		52, 54, 5, 7, 0, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58,
		60, 3, 12, 6, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0,
		0, 0, 61, 63, 5, 8, 0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 5,
		1, 0, 0, 0, 64, 65, 7, 0, 0, 0, 65, 7, 1, 0, 0, 0, 66, 67, 7, 1, 0, 0,
		67, 9, 1, 0, 0, 0, 68, 70, 5, 2, 0, 0, 69, 71, 5, 8, 0, 0, 70, 69, 1, 0,
		0, 0, 70, 71, 1, 0, 0, 0, 71, 11, 1, 0, 0, 0, 72, 73, 5, 2, 0, 0, 73, 13,
		1, 0, 0, 0, 11, 17, 25, 30, 37, 44, 48, 50, 55, 59, 62, 70,
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

// EnvLangFileParserInit initializes any static state used to implement EnvLangFileParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEnvLangFileParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EnvLangFileParserInit() {
	staticData := &EnvLangFileParserStaticData
	staticData.once.Do(envlangfileParserInit)
}

// NewEnvLangFileParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEnvLangFileParser(input antlr.TokenStream) *EnvLangFileParser {
	EnvLangFileParserInit()
	this := new(EnvLangFileParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EnvLangFileParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EnvLangFile.g4"

	return this
}

// EnvLangFileParser tokens.
const (
	EnvLangFileParserEOF            = antlr.TokenEOF
	EnvLangFileParserASSIGN         = 1
	EnvLangFileParserCOMMENT        = 2
	EnvLangFileParserDQSTRING       = 3
	EnvLangFileParserSQSTRING       = 4
	EnvLangFileParserUNQUOTED_VALUE = 5
	EnvLangFileParserIDENTIFIER     = 6
	EnvLangFileParserWS             = 7
	EnvLangFileParserNEWLINE        = 8
)

// EnvLangFileParser rules.
const (
	EnvLangFileParserRULE_envFile       = 0
	EnvLangFileParserRULE_line          = 1
	EnvLangFileParserRULE_entry         = 2
	EnvLangFileParserRULE_identifier    = 3
	EnvLangFileParserRULE_value         = 4
	EnvLangFileParserRULE_comment       = 5
	EnvLangFileParserRULE_inlineComment = 6
)

// IEnvFileContext is an interface to support dynamic dispatch.
type IEnvFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext

	// IsEnvFileContext differentiates from other interfaces.
	IsEnvFileContext()
}

type EnvFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvFileContext() *EnvFileContext {
	var p = new(EnvFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_envFile
	return p
}

func InitEmptyEnvFileContext(p *EnvFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_envFile
}

func (*EnvFileContext) IsEnvFileContext() {}

func NewEnvFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvFileContext {
	var p = new(EnvFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_envFile

	return p
}

func (s *EnvFileContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserEOF, 0)
}

func (s *EnvFileContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *EnvFileContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *EnvFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterEnvFile(s)
	}
}

func (s *EnvFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitEnvFile(s)
	}
}

func (p *EnvLangFileParser) EnvFile() (localctx IEnvFileContext) {
	localctx = NewEnvFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EnvLangFileParserRULE_envFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&484) != 0 {
		{
			p.SetState(14)
			p.Line()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(EnvLangFileParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Entry() IEntryContext
	Comment() ICommentContext
	NEWLINE() antlr.TerminalNode

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Entry() IEntryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserNEWLINE, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *EnvLangFileParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EnvLangFileParserRULE_line)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EnvLangFileParserUNQUOTED_VALUE, EnvLangFileParserIDENTIFIER, EnvLangFileParserWS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Entry()
		}

	case EnvLangFileParserCOMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Comment()
		}

	case EnvLangFileParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(EnvLangFileParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	InlineComment() IInlineCommentContext
	NEWLINE() antlr.TerminalNode
	Value() IValueContext

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EntryContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(EnvLangFileParserWS)
}

func (s *EntryContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserWS, i)
}

func (s *EntryContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserASSIGN, 0)
}

func (s *EntryContext) InlineComment() IInlineCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineCommentContext)
}

func (s *EntryContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserNEWLINE, 0)
}

func (s *EntryContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *EnvLangFileParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EnvLangFileParserRULE_entry)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangFileParserWS {
		{
			p.SetState(27)
			p.Match(EnvLangFileParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Identifier()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(34)
				p.Match(EnvLangFileParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EnvLangFileParserASSIGN {
		{
			p.SetState(40)
			p.Match(EnvLangFileParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(41)
					p.Match(EnvLangFileParserWS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(47)
				p.Value()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(52)
				p.Match(EnvLangFileParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(58)
			p.InlineComment()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(61)
			p.Match(EnvLangFileParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	UNQUOTED_VALUE() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserIDENTIFIER, 0)
}

func (s *IdentifierContext) UNQUOTED_VALUE() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserUNQUOTED_VALUE, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *EnvLangFileParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangFileParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EnvLangFileParserUNQUOTED_VALUE || _la == EnvLangFileParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DQSTRING() antlr.TerminalNode
	SQSTRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	UNQUOTED_VALUE() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DQSTRING() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserDQSTRING, 0)
}

func (s *ValueContext) SQSTRING() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserSQSTRING, 0)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserIDENTIFIER, 0)
}

func (s *ValueContext) UNQUOTED_VALUE() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserUNQUOTED_VALUE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *EnvLangFileParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EnvLangFileParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserCOMMENT, 0)
}

func (s *CommentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserNEWLINE, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *EnvLangFileParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EnvLangFileParserRULE_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(EnvLangFileParserCOMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(69)
			p.Match(EnvLangFileParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInlineCommentContext is an interface to support dynamic dispatch.
type IInlineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsInlineCommentContext differentiates from other interfaces.
	IsInlineCommentContext()
}

type InlineCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineCommentContext() *InlineCommentContext {
	var p = new(InlineCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_inlineComment
	return p
}

func InitEmptyInlineCommentContext(p *InlineCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_inlineComment
}

func (*InlineCommentContext) IsInlineCommentContext() {}

func NewInlineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineCommentContext {
	var p = new(InlineCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_inlineComment

	return p
}

func (s *InlineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineCommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserCOMMENT, 0)
}

func (s *InlineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterInlineComment(s)
	}
}

func (s *InlineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitInlineComment(s)
	}
}

func (p *EnvLangFileParser) InlineComment() (localctx IInlineCommentContext) {
	localctx = NewInlineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EnvLangFileParserRULE_inlineComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(EnvLangFileParserCOMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
