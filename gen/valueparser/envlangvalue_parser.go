// Code generated from EnvLangValue.g4 by ANTLR 4.9.3. DO NOT EDIT.

package valueparser // EnvLangValue
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 71, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	2, 7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 5, 2, 32, 10, 2, 3,
	3, 3, 3, 7, 3, 36, 10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 7, 4, 45, 10, 4, 12, 4, 14, 4, 48, 11, 4, 3, 5, 3, 5, 5, 5, 52, 10,
	5, 3, 6, 3, 6, 7, 6, 56, 10, 6, 12, 6, 14, 6, 59, 11, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 69, 10, 9, 3, 9, 2, 2, 10, 2, 4,
	6, 8, 10, 12, 14, 16, 2, 2, 2, 72, 2, 21, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2,
	6, 42, 3, 2, 2, 2, 8, 51, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 60, 3, 2,
	2, 2, 14, 62, 3, 2, 2, 2, 16, 68, 3, 2, 2, 2, 18, 20, 5, 16, 9, 2, 19,
	18, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2,
	2, 22, 31, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 26, 7, 10, 2, 2, 25, 24,
	3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2,
	28, 32, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32, 7, 2, 2, 3, 31, 27, 3,
	2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 3, 3, 2, 2, 2, 33, 37, 7, 3, 2, 2, 34,
	36, 7, 8, 2, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2,
	2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 41,
	7, 4, 2, 2, 41, 5, 3, 2, 2, 2, 42, 46, 7, 5, 2, 2, 43, 45, 7, 8, 2, 2,
	44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 52, 5, 4, 3, 2, 50,
	52, 5, 6, 4, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 9, 3, 2, 2,
	2, 53, 57, 7, 8, 2, 2, 54, 56, 7, 9, 2, 2, 55, 54, 3, 2, 2, 2, 56, 59,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 11, 3, 2, 2, 2,
	59, 57, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2, 61, 13, 3, 2, 2, 2, 62, 63, 7,
	7, 2, 2, 63, 15, 3, 2, 2, 2, 64, 69, 5, 14, 8, 2, 65, 69, 5, 8, 5, 2, 66,
	69, 5, 12, 7, 2, 67, 69, 5, 10, 6, 2, 68, 64, 3, 2, 2, 2, 68, 65, 3, 2,
	2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 17, 3, 2, 2, 2, 10, 21,
	27, 31, 37, 46, 51, 57, 68,
}
var literalNames = []string{
	"", "'${'", "'}'", "'$'", "' '", "'\"\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "TEXT_NO_SPACE", "TEXT_ANY", "CRLF",
}

var ruleNames = []string{
	"dqstring", "strictVar", "simpleVar", "variable", "text", "space", "dQEscape",
	"content",
}

type EnvLangValueParser struct {
	*antlr.BaseParser
}

// NewEnvLangValueParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *EnvLangValueParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEnvLangValueParser(input antlr.TokenStream) *EnvLangValueParser {
	this := new(EnvLangValueParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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
	EnvLangValueParserTEXT_NO_SPACE = 6
	EnvLangValueParserTEXT_ANY      = 7
	EnvLangValueParserCRLF          = 8
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
	EnvLangValueParserRULE_content   = 7
)

// IDqstringContext is an interface to support dynamic dispatch.
type IDqstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDqstringContext differentiates from other interfaces.
	IsDqstringContext()
}

type DqstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDqstringContext() *DqstringContext {
	var p = new(DqstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dqstring
	return p
}

func (*DqstringContext) IsDqstringContext() {}

func NewDqstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DqstringContext {
	var p = new(DqstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_dqstring

	return p
}

func (s *DqstringContext) GetParser() antlr.Parser { return s.parser }

func (s *DqstringContext) EOF() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserEOF, 0)
}

func (s *DqstringContext) AllContent() []IContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContentContext)(nil)).Elem())
	var tst = make([]IContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContentContext)
		}
	}

	return tst
}

func (s *DqstringContext) Content(i int) IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), i)

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
	this := p
	_ = this

	localctx = NewDqstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EnvLangValueParserRULE_dqstring)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EnvLangValueParserT__0)|(1<<EnvLangValueParserT__2)|(1<<EnvLangValueParserT__3)|(1<<EnvLangValueParserT__4)|(1<<EnvLangValueParserTEXT_NO_SPACE))) != 0 {
		{
			p.SetState(16)
			p.Content()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangValueParserCRLF {
			{
				p.SetState(22)
				p.Match(EnvLangValueParserCRLF)
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(28)
			p.Match(EnvLangValueParserEOF)
		}

	}

	return localctx
}

// IStrictVarContext is an interface to support dynamic dispatch.
type IStrictVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrictVarContext differentiates from other interfaces.
	IsStrictVarContext()
}

type StrictVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrictVarContext() *StrictVarContext {
	var p = new(StrictVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_strictVar
	return p
}

func (*StrictVarContext) IsStrictVarContext() {}

func NewStrictVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrictVarContext {
	var p = new(StrictVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewStrictVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EnvLangValueParserRULE_strictVar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(EnvLangValueParserT__0)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangValueParserTEXT_NO_SPACE {
		{
			p.SetState(32)
			p.Match(EnvLangValueParserTEXT_NO_SPACE)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(EnvLangValueParserT__1)
	}

	return localctx
}

// ISimpleVarContext is an interface to support dynamic dispatch.
type ISimpleVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleVarContext differentiates from other interfaces.
	IsSimpleVarContext()
}

type SimpleVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleVarContext() *SimpleVarContext {
	var p = new(SimpleVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_simpleVar
	return p
}

func (*SimpleVarContext) IsSimpleVarContext() {}

func NewSimpleVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleVarContext {
	var p = new(SimpleVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSimpleVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EnvLangValueParserRULE_simpleVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(EnvLangValueParserT__2)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.Match(EnvLangValueParserTEXT_NO_SPACE)
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) StrictVar() IStrictVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrictVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrictVarContext)
}

func (s *VariableContext) SimpleVar() ISimpleVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleVarContext)(nil)).Elem(), 0)

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
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangValueParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EnvLangValueParserRULE_text)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(EnvLangValueParserTEXT_NO_SPACE)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangValueParserTEXT_ANY {
		{
			p.SetState(52)
			p.Match(EnvLangValueParserTEXT_ANY)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpaceContext() *SpaceContext {
	var p = new(SpaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_space
	return p
}

func (*SpaceContext) IsSpaceContext() {}

func NewSpaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpaceContext {
	var p = new(SpaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSpaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EnvLangValueParserRULE_space)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(EnvLangValueParserT__3)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDQEscapeContext() *DQEscapeContext {
	var p = new(DQEscapeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_dQEscape
	return p
}

func (*DQEscapeContext) IsDQEscapeContext() {}

func NewDQEscapeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DQEscapeContext {
	var p = new(DQEscapeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewDQEscapeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EnvLangValueParserRULE_dQEscape)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(EnvLangValueParserT__4)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) DQEscape() IDQEscapeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDQEscapeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDQEscapeContext)
}

func (s *ContentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ContentContext) Space() ISpaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpaceContext)
}

func (s *ContentContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

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
	this := p
	_ = this

	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EnvLangValueParserRULE_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EnvLangValueParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.DQEscape()
		}

	case EnvLangValueParserT__0, EnvLangValueParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Variable()
		}

	case EnvLangValueParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Space()
		}

	case EnvLangValueParserTEXT_NO_SPACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Text()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
