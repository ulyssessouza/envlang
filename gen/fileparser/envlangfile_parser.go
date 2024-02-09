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
		"", "'='", "'\\r'", "'\\n'", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"envFile", "entry", "identifier", "assign", "value", "nl",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 57, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0, 5,
		0, 21, 8, 0, 10, 0, 12, 0, 24, 9, 0, 4, 0, 26, 8, 0, 11, 0, 12, 0, 27,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 1, 5, 1, 38, 8, 1, 10,
		1, 12, 1, 41, 9, 1, 1, 1, 3, 1, 44, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 3, 5, 53, 8, 5, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10,
		0, 1, 1, 0, 4, 6, 58, 0, 15, 1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 4, 45, 1, 0,
		0, 0, 6, 47, 1, 0, 0, 0, 8, 49, 1, 0, 0, 0, 10, 52, 1, 0, 0, 0, 12, 14,
		3, 10, 5, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0,
		15, 16, 1, 0, 0, 0, 16, 25, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 22, 3,
		2, 1, 0, 19, 21, 3, 10, 5, 0, 20, 19, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22,
		20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0,
		0, 25, 18, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28,
		1, 0, 0, 0, 28, 1, 1, 0, 0, 0, 29, 34, 3, 4, 2, 0, 30, 31, 3, 6, 3, 0,
		31, 32, 3, 8, 4, 0, 32, 35, 1, 0, 0, 0, 33, 35, 3, 6, 3, 0, 34, 30, 1,
		0, 0, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 43, 1, 0, 0, 0, 36,
		38, 3, 10, 5, 0, 37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0,
		0, 0, 39, 40, 1, 0, 0, 0, 40, 44, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 44,
		5, 0, 0, 1, 43, 39, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 3, 1, 0, 0, 0,
		45, 46, 5, 4, 0, 0, 46, 5, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 7, 1, 0,
		0, 0, 49, 50, 7, 0, 0, 0, 50, 9, 1, 0, 0, 0, 51, 53, 5, 2, 0, 0, 52, 51,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 5, 3, 0, 0,
		55, 11, 1, 0, 0, 0, 7, 15, 22, 27, 34, 39, 43, 52,
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
	EnvLangFileParserEOF      = antlr.TokenEOF
	EnvLangFileParserT__0     = 1
	EnvLangFileParserT__1     = 2
	EnvLangFileParserT__2     = 3
	EnvLangFileParserTEXT     = 4
	EnvLangFileParserDQSTRING = 5
	EnvLangFileParserSQSTRING = 6
	EnvLangFileParserSPACE    = 7
)

// EnvLangFileParser rules.
const (
	EnvLangFileParserRULE_envFile    = 0
	EnvLangFileParserRULE_entry      = 1
	EnvLangFileParserRULE_identifier = 2
	EnvLangFileParserRULE_assign     = 3
	EnvLangFileParserRULE_value      = 4
	EnvLangFileParserRULE_nl         = 5
)

// IEnvFileContext is an interface to support dynamic dispatch.
type IEnvFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNl() []INlContext
	Nl(i int) INlContext
	AllEntry() []IEntryContext
	Entry(i int) IEntryContext

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

func (s *EnvFileContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *EnvFileContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
}

func (s *EnvFileContext) AllEntry() []IEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEntryContext); ok {
			len++
		}
	}

	tst := make([]IEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEntryContext); ok {
			tst[i] = t.(IEntryContext)
			i++
		}
	}

	return tst
}

func (s *EnvFileContext) Entry(i int) IEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryContext); ok {
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

	return t.(IEntryContext)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangFileParserT__1 || _la == EnvLangFileParserT__2 {
		{
			p.SetState(12)
			p.Nl()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EnvLangFileParserTEXT {
		{
			p.SetState(18)
			p.Entry()
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangFileParserT__1 || _la == EnvLangFileParserT__2 {
			{
				p.SetState(19)
				p.Nl()
			}

			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	EOF() antlr.TerminalNode
	Assign() IAssignContext
	Value() IValueContext
	AllNl() []INlContext
	Nl(i int) INlContext

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

func (s *EntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserEOF, 0)
}

func (s *EntryContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
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

func (s *EntryContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *EntryContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
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
	p.EnterRule(localctx, 2, EnvLangFileParserRULE_entry)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Identifier()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(30)
			p.Assign()
		}
		{
			p.SetState(31)
			p.Value()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(33)
			p.Assign()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.SetState(39)
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
					p.SetState(36)
					p.Nl()
				}

			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(42)
			p.Match(EnvLangFileParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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
	TEXT() antlr.TerminalNode

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

func (s *IdentifierContext) TEXT() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserTEXT, 0)
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
	p.EnterRule(localctx, 4, EnvLangFileParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(EnvLangFileParserTEXT)
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

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *EnvLangFileParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangFileParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(EnvLangFileParserT__0)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	DQSTRING() antlr.TerminalNode
	SQSTRING() antlr.TerminalNode
	TEXT() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *ValueContext) GetOp() antlr.Token { return s.op }

func (s *ValueContext) SetOp(v antlr.Token) { s.op = v }

func (s *ValueContext) DQSTRING() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserDQSTRING, 0)
}

func (s *ValueContext) SQSTRING() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserSQSTRING, 0)
}

func (s *ValueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserTEXT, 0)
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
		p.SetState(49)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ValueContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ValueContext).op = _ri
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

// INlContext is an interface to support dynamic dispatch.
type INlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNlContext differentiates from other interfaces.
	IsNlContext()
}

type NlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNlContext() *NlContext {
	var p = new(NlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_nl
	return p
}

func InitEmptyNlContext(p *NlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_nl
}

func (*NlContext) IsNlContext() {}

func NewNlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NlContext {
	var p = new(NlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_nl

	return p
}

func (s *NlContext) GetParser() antlr.Parser { return s.parser }
func (s *NlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.EnterNl(s)
	}
}

func (s *NlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangFileListener); ok {
		listenerT.ExitNl(s)
	}
}

func (p *EnvLangFileParser) Nl() (localctx INlContext) {
	localctx = NewNlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EnvLangFileParserRULE_nl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EnvLangFileParserT__1 {
		{
			p.SetState(51)
			p.Match(EnvLangFileParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(54)
		p.Match(EnvLangFileParserT__2)
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
