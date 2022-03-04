// Code generated from EnvLangFile.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fileparser // EnvLangFile
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 59, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 7, 2, 23, 10,
	2, 12, 2, 14, 2, 26, 11, 2, 6, 2, 28, 10, 2, 13, 2, 14, 2, 29, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 37, 10, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14,
	3, 43, 11, 3, 3, 3, 5, 3, 46, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 5, 7, 55, 10, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2,
	3, 3, 2, 6, 8, 2, 60, 2, 17, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6, 47, 3, 2,
	2, 2, 8, 49, 3, 2, 2, 2, 10, 51, 3, 2, 2, 2, 12, 54, 3, 2, 2, 2, 14, 16,
	5, 12, 7, 2, 15, 14, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2,
	17, 18, 3, 2, 2, 2, 18, 27, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 24, 5,
	4, 3, 2, 21, 23, 5, 12, 7, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24,
	22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2,
	2, 27, 20, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30,
	3, 2, 2, 2, 30, 3, 3, 2, 2, 2, 31, 36, 5, 6, 4, 2, 32, 33, 5, 8, 5, 2,
	33, 34, 5, 10, 6, 2, 34, 37, 3, 2, 2, 2, 35, 37, 5, 8, 5, 2, 36, 32, 3,
	2, 2, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 45, 3, 2, 2, 2, 38,
	40, 5, 12, 7, 2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2,
	2, 2, 41, 42, 3, 2, 2, 2, 42, 46, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 46,
	7, 2, 2, 3, 45, 41, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 5, 3, 2, 2, 2,
	47, 48, 7, 6, 2, 2, 48, 7, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2, 50, 9, 3, 2,
	2, 2, 51, 52, 9, 2, 2, 2, 52, 11, 3, 2, 2, 2, 53, 55, 7, 4, 2, 2, 54, 53,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 7, 5, 2, 2,
	57, 13, 3, 2, 2, 2, 9, 17, 24, 29, 36, 41, 45, 54,
}
var literalNames = []string{
	"", "'='", "'\r'", "'\n'", "", "", "", "' '",
}
var symbolicNames = []string{
	"", "", "", "", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
}

var ruleNames = []string{
	"envFile", "entry", "identifier", "assign", "value", "nl",
}

type EnvLangFileParser struct {
	*antlr.BaseParser
}

// NewEnvLangFileParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *EnvLangFileParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEnvLangFileParser(input antlr.TokenStream) *EnvLangFileParser {
	this := new(EnvLangFileParser)
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

	// IsEnvFileContext differentiates from other interfaces.
	IsEnvFileContext()
}

type EnvFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvFileContext() *EnvFileContext {
	var p = new(EnvFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_envFile
	return p
}

func (*EnvFileContext) IsEnvFileContext() {}

func NewEnvFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvFileContext {
	var p = new(EnvFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_envFile

	return p
}

func (s *EnvFileContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvFileContext) AllNl() []INlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INlContext)(nil)).Elem())
	var tst = make([]INlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INlContext)
		}
	}

	return tst
}

func (s *EnvFileContext) Nl(i int) INlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *EnvFileContext) AllEntry() []IEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntryContext)(nil)).Elem())
	var tst = make([]IEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntryContext)
		}
	}

	return tst
}

func (s *EnvFileContext) Entry(i int) IEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntryContext)(nil)).Elem(), i)

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
	this := p
	_ = this

	localctx = NewEnvFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EnvLangFileParserRULE_envFile)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangFileParserT__1 || _la == EnvLangFileParserT__2 {
		{
			p.SetState(12)
			p.Nl()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EnvLangFileParserTEXT {
		{
			p.SetState(18)
			p.Entry()
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangFileParserT__1 || _la == EnvLangFileParserT__2 {
			{
				p.SetState(19)
				p.Nl()
			}

			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_entry
	return p
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangFileParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserEOF, 0)
}

func (s *EntryContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *EntryContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EntryContext) AllNl() []INlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INlContext)(nil)).Elem())
	var tst = make([]INlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INlContext)
		}
	}

	return tst
}

func (s *EntryContext) Nl(i int) INlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INlContext)(nil)).Elem(), i)

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
	this := p
	_ = this

	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EnvLangFileParserRULE_entry)

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
		p.SetState(29)
		p.Identifier()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(30)
			p.Assign()
		}
		{
			p.SetState(31)
			p.Value()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(33)
			p.Assign()
		}

	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(36)
					p.Nl()
				}

			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case 2:
		{
			p.SetState(42)
			p.Match(EnvLangFileParserEOF)
		}

	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EnvLangFileParserRULE_identifier)

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
		p.SetState(45)
		p.Match(EnvLangFileParserTEXT)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangFileParserRULE_assign)

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
		p.SetState(47)
		p.Match(EnvLangFileParserT__0)
	}

	return localctx
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

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EnvLangFileParserRULE_value)
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
		p.SetState(49)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ValueContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EnvLangFileParserTEXT)|(1<<EnvLangFileParserDQSTRING)|(1<<EnvLangFileParserSQSTRING))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ValueContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNlContext() *NlContext {
	var p = new(NlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EnvLangFileParserRULE_nl
	return p
}

func (*NlContext) IsNlContext() {}

func NewNlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NlContext {
	var p = new(NlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewNlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EnvLangFileParserRULE_nl)
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EnvLangFileParserT__1 {
		{
			p.SetState(51)
			p.Match(EnvLangFileParserT__1)
		}

	}
	{
		p.SetState(54)
		p.Match(EnvLangFileParserT__2)
	}

	return localctx
}
