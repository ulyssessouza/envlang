// Code generated from EnvLangValue.g4 by ANTLR 4.13.1. DO NOT EDIT.

package valueparser // EnvLangValue
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

type EnvLangValueParser struct {
	*antlr.BaseParser
}

var EnvLangValueParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func envlangvalueParserInit() {
	staticData := &EnvLangValueParserStaticData
	staticData.LiteralNames = []string{
		"", "'${'", "'}'", "'$'", "' '", "'\"\"'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
	}
	staticData.RuleNames = []string{
		"dqstring", "strictVar", "simpleVar", "variable", "text", "space", "dQEscape",
		"special", "content",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 72, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 5, 0, 20, 8, 0, 10,
		0, 12, 0, 23, 9, 0, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0,
		3, 0, 32, 8, 0, 1, 1, 1, 1, 4, 1, 36, 8, 1, 11, 1, 12, 1, 37, 1, 1, 1,
		1, 1, 2, 1, 2, 4, 2, 44, 8, 2, 11, 2, 12, 2, 45, 1, 3, 1, 3, 3, 3, 50,
		8, 3, 1, 4, 1, 4, 5, 4, 54, 8, 4, 10, 4, 12, 4, 57, 9, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 70, 8, 8, 1, 8,
		0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 2, 0, 2, 3, 6, 6, 73, 0,
		21, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4, 41, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0,
		8, 51, 1, 0, 0, 0, 10, 58, 1, 0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 62, 1, 0,
		0, 0, 16, 69, 1, 0, 0, 0, 18, 20, 3, 16, 8, 0, 19, 18, 1, 0, 0, 0, 20,
		23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 31, 1, 0, 0,
		0, 23, 21, 1, 0, 0, 0, 24, 26, 5, 9, 0, 0, 25, 24, 1, 0, 0, 0, 26, 29,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 32, 1, 0, 0, 0,
		29, 27, 1, 0, 0, 0, 30, 32, 5, 0, 0, 1, 31, 27, 1, 0, 0, 0, 31, 30, 1,
		0, 0, 0, 32, 1, 1, 0, 0, 0, 33, 35, 5, 1, 0, 0, 34, 36, 5, 7, 0, 0, 35,
		34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0,
		0, 38, 39, 1, 0, 0, 0, 39, 40, 5, 2, 0, 0, 40, 3, 1, 0, 0, 0, 41, 43, 5,
		3, 0, 0, 42, 44, 5, 7, 0, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45,
		43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 5, 1, 0, 0, 0, 47, 50, 3, 2, 1,
		0, 48, 50, 3, 4, 2, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 7, 1,
		0, 0, 0, 51, 55, 5, 7, 0, 0, 52, 54, 5, 8, 0, 0, 53, 52, 1, 0, 0, 0, 54,
		57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 9, 1, 0, 0,
		0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 4, 0, 0, 59, 11, 1, 0, 0, 0, 60, 61,
		5, 5, 0, 0, 61, 13, 1, 0, 0, 0, 62, 63, 7, 0, 0, 0, 63, 15, 1, 0, 0, 0,
		64, 70, 3, 12, 6, 0, 65, 70, 3, 6, 3, 0, 66, 70, 3, 10, 5, 0, 67, 70, 3,
		14, 7, 0, 68, 70, 3, 8, 4, 0, 69, 64, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 69,
		66, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 17, 1, 0, 0,
		0, 8, 21, 27, 31, 37, 45, 49, 55, 69,
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

// EnvLangValueParserInit initializes any static state used to implement EnvLangValueParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEnvLangValueParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EnvLangValueParserInit() {
	staticData := &EnvLangValueParserStaticData
	staticData.once.Do(envlangvalueParserInit)
}

// NewEnvLangValueParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEnvLangValueParser(input antlr.TokenStream) *EnvLangValueParser {
	EnvLangValueParserInit()
	this := new(EnvLangValueParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EnvLangValueParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EnvLangValue.g4"

	return this
}

// EnvLangValueParser tokens.
const (
	EnvLangValueParserEOF           = antlr.TokenEOF
	EnvLangValueParserT__0          = 1
	EnvLangValueParserT__1          = 2
	EnvLangValueParserT__2          = 3
	EnvLangValueParserT__3          = 4
	EnvLangValueParserT__4          = 5
	EnvLangValueParserT__5          = 6
	EnvLangValueParserTEXT_NO_SPACE = 7
	EnvLangValueParserTEXT_ANY      = 8
	EnvLangValueParserCRLF          = 9
)

// EnvLangValueParser rules.
const (
	EnvLangValueParserRULE_dqstring  = 0
	EnvLangValueParserRULE_strictVar = 1
	EnvLangValueParserRULE_simpleVar = 2
	EnvLangValueParserRULE_variable  = 3
	EnvLangValueParserRULE_text      = 4
	EnvLangValueParserRULE_space     = 5
	EnvLangValueParserRULE_dQEscape  = 6
	EnvLangValueParserRULE_special   = 7
	EnvLangValueParserRULE_content   = 8
)

// IDqstringContext is an interface to support dynamic dispatch.
type IDqstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllContent() []IContentContext
	Content(i int) IContentContext
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode

	// IsDqstringContext differentiates from other interfaces.
	IsDqstringContext()
}

type DqstringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDqstringContext() *DqstringContext {
	var p = new(DqstringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dqstring
	return p
}

func InitEmptyDqstringContext(p *DqstringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dqstring
}

func (*DqstringContext) IsDqstringContext() {}

func NewDqstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DqstringContext {
	var p = new(DqstringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_dqstring

	return p
}

func (s *DqstringContext) GetParser() antlr.Parser { return s.parser }

func (s *DqstringContext) EOF() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserEOF, 0)
}

func (s *DqstringContext) AllContent() []IContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentContext); ok {
			len++
		}
	}

	tst := make([]IContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentContext); ok {
			tst[i] = t.(IContentContext)
			i++
		}
	}

	return tst
}

func (s *DqstringContext) Content(i int) IContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentContext); ok {
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

	return t.(IContentContext)
}

func (s *DqstringContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(EnvLangValueParserCRLF)
}

func (s *DqstringContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserCRLF, i)
}

func (s *DqstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DqstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DqstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterDqstring(s)
	}
}

func (s *DqstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitDqstring(s)
	}
}

func (p *EnvLangValueParser) Dqstring() (localctx IDqstringContext) {
	localctx = NewDqstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EnvLangValueParserRULE_dqstring)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254) != 0 {
		{
			p.SetState(18)
			p.Content()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangValueParserCRLF {
			{
				p.SetState(24)
				p.Match(EnvLangValueParserCRLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(30)
			p.Match(EnvLangValueParserEOF)
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

// IStrictVarContext is an interface to support dynamic dispatch.
type IStrictVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT_NO_SPACE() []antlr.TerminalNode
	TEXT_NO_SPACE(i int) antlr.TerminalNode

	// IsStrictVarContext differentiates from other interfaces.
	IsStrictVarContext()
}

type StrictVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrictVarContext() *StrictVarContext {
	var p = new(StrictVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_strictVar
	return p
}

func InitEmptyStrictVarContext(p *StrictVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_strictVar
}

func (*StrictVarContext) IsStrictVarContext() {}

func NewStrictVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrictVarContext {
	var p = new(StrictVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_strictVar

	return p
}

func (s *StrictVarContext) GetParser() antlr.Parser { return s.parser }

func (s *StrictVarContext) AllTEXT_NO_SPACE() []antlr.TerminalNode {
	return s.GetTokens(EnvLangValueParserTEXT_NO_SPACE)
}

func (s *StrictVarContext) TEXT_NO_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserTEXT_NO_SPACE, i)
}

func (s *StrictVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrictVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrictVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterStrictVar(s)
	}
}

func (s *StrictVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitStrictVar(s)
	}
}

func (p *EnvLangValueParser) StrictVar() (localctx IStrictVarContext) {
	localctx = NewStrictVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EnvLangValueParserRULE_strictVar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(EnvLangValueParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EnvLangValueParserTEXT_NO_SPACE {
		{
			p.SetState(34)
			p.Match(EnvLangValueParserTEXT_NO_SPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(EnvLangValueParserT__1)
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

// ISimpleVarContext is an interface to support dynamic dispatch.
type ISimpleVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT_NO_SPACE() []antlr.TerminalNode
	TEXT_NO_SPACE(i int) antlr.TerminalNode

	// IsSimpleVarContext differentiates from other interfaces.
	IsSimpleVarContext()
}

type SimpleVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleVarContext() *SimpleVarContext {
	var p = new(SimpleVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_simpleVar
	return p
}

func InitEmptySimpleVarContext(p *SimpleVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_simpleVar
}

func (*SimpleVarContext) IsSimpleVarContext() {}

func NewSimpleVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleVarContext {
	var p = new(SimpleVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_simpleVar

	return p
}

func (s *SimpleVarContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleVarContext) AllTEXT_NO_SPACE() []antlr.TerminalNode {
	return s.GetTokens(EnvLangValueParserTEXT_NO_SPACE)
}

func (s *SimpleVarContext) TEXT_NO_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserTEXT_NO_SPACE, i)
}

func (s *SimpleVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterSimpleVar(s)
	}
}

func (s *SimpleVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitSimpleVar(s)
	}
}

func (p *EnvLangValueParser) SimpleVar() (localctx ISimpleVarContext) {
	localctx = NewSimpleVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EnvLangValueParserRULE_simpleVar)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(EnvLangValueParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(42)
				p.Match(EnvLangValueParserTEXT_NO_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StrictVar() IStrictVarContext
	SimpleVar() ISimpleVarContext

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) StrictVar() IStrictVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrictVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrictVarContext)
}

func (s *VariableContext) SimpleVar() ISimpleVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleVarContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *EnvLangValueParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangValueParserRULE_variable)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EnvLangValueParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.StrictVar()
		}

	case EnvLangValueParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.SimpleVar()
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

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_NO_SPACE() antlr.TerminalNode
	AllTEXT_ANY() []antlr.TerminalNode
	TEXT_ANY(i int) antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TEXT_NO_SPACE() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserTEXT_NO_SPACE, 0)
}

func (s *TextContext) AllTEXT_ANY() []antlr.TerminalNode {
	return s.GetTokens(EnvLangValueParserTEXT_ANY)
}

func (s *TextContext) TEXT_ANY(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserTEXT_ANY, i)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *EnvLangValueParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EnvLangValueParserRULE_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(EnvLangValueParserTEXT_NO_SPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangValueParserTEXT_ANY {
		{
			p.SetState(52)
			p.Match(EnvLangValueParserTEXT_ANY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(57)
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

// ISpaceContext is an interface to support dynamic dispatch.
type ISpaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSpaceContext differentiates from other interfaces.
	IsSpaceContext()
}

type SpaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpaceContext() *SpaceContext {
	var p = new(SpaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_space
	return p
}

func InitEmptySpaceContext(p *SpaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_space
}

func (*SpaceContext) IsSpaceContext() {}

func NewSpaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpaceContext {
	var p = new(SpaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_space

	return p
}

func (s *SpaceContext) GetParser() antlr.Parser { return s.parser }
func (s *SpaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterSpace(s)
	}
}

func (s *SpaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitSpace(s)
	}
}

func (p *EnvLangValueParser) Space() (localctx ISpaceContext) {
	localctx = NewSpaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EnvLangValueParserRULE_space)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(EnvLangValueParserT__3)
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

// IDQEscapeContext is an interface to support dynamic dispatch.
type IDQEscapeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDQEscapeContext differentiates from other interfaces.
	IsDQEscapeContext()
}

type DQEscapeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDQEscapeContext() *DQEscapeContext {
	var p = new(DQEscapeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dQEscape
	return p
}

func InitEmptyDQEscapeContext(p *DQEscapeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dQEscape
}

func (*DQEscapeContext) IsDQEscapeContext() {}

func NewDQEscapeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DQEscapeContext {
	var p = new(DQEscapeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_dQEscape

	return p
}

func (s *DQEscapeContext) GetParser() antlr.Parser { return s.parser }
func (s *DQEscapeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DQEscapeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DQEscapeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterDQEscape(s)
	}
}

func (s *DQEscapeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitDQEscape(s)
	}
}

func (p *EnvLangValueParser) DQEscape() (localctx IDQEscapeContext) {
	localctx = NewDQEscapeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EnvLangValueParserRULE_dQEscape)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(EnvLangValueParserT__4)
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

// ISpecialContext is an interface to support dynamic dispatch.
type ISpecialContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSpecialContext differentiates from other interfaces.
	IsSpecialContext()
}

type SpecialContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialContext() *SpecialContext {
	var p = new(SpecialContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_special
	return p
}

func InitEmptySpecialContext(p *SpecialContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_special
}

func (*SpecialContext) IsSpecialContext() {}

func NewSpecialContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialContext {
	var p = new(SpecialContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_special

	return p
}

func (s *SpecialContext) GetParser() antlr.Parser { return s.parser }
func (s *SpecialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterSpecial(s)
	}
}

func (s *SpecialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitSpecial(s)
	}
}

func (p *EnvLangValueParser) Special() (localctx ISpecialContext) {
	localctx = NewSpecialContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EnvLangValueParserRULE_special)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&76) != 0) {
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

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DQEscape() IDQEscapeContext
	Variable() IVariableContext
	Space() ISpaceContext
	Special() ISpecialContext
	Text() ITextContext

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_content
	return p
}

func InitEmptyContentContext(p *ContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_content
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) DQEscape() IDQEscapeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDQEscapeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDQEscapeContext)
}

func (s *ContentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ContentContext) Space() ISpaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpaceContext)
}

func (s *ContentContext) Special() ISpecialContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialContext)
}

func (s *ContentContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *EnvLangValueParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EnvLangValueParserRULE_content)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.DQEscape()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Space()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Special()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.Text()
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
