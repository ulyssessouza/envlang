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
	staticData.SymbolicNames = []string{
		"", "STRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_ASSIGN_IF_UNSET",
		"STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "DOLLAR", "ESCAPED_CHAR", "TEXT",
		"WS", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"dqstring", "content", "variable", "escapedChar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 28, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 22, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6,
		0, 1, 1, 0, 1, 7, 28, 0, 11, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 23, 1, 0,
		0, 0, 6, 25, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1,
		0, 0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13,
		11, 1, 0, 0, 0, 14, 15, 5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 22, 3, 4, 2,
		0, 17, 22, 3, 6, 3, 0, 18, 22, 5, 9, 0, 0, 19, 22, 5, 10, 0, 0, 20, 22,
		5, 11, 0, 0, 21, 16, 1, 0, 0, 0, 21, 17, 1, 0, 0, 0, 21, 18, 1, 0, 0, 0,
		21, 19, 1, 0, 0, 0, 21, 20, 1, 0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 24, 7, 0,
		0, 0, 24, 5, 1, 0, 0, 0, 25, 26, 5, 8, 0, 0, 26, 7, 1, 0, 0, 0, 2, 11,
		21,
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
	EnvLangValueParserEOF                                       = antlr.TokenEOF
	EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY  = 1
	EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET           = 2
	EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY = 3
	EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET          = 4
	EnvLangValueParserSIMPLE_STRICT_VAR                         = 5
	EnvLangValueParserSIMPLE_VAR                                = 6
	EnvLangValueParserDOLLAR                                    = 7
	EnvLangValueParserESCAPED_CHAR                              = 8
	EnvLangValueParserTEXT                                      = 9
	EnvLangValueParserWS                                        = 10
	EnvLangValueParserNEWLINE                                   = 11
)

// EnvLangValueParser rules.
const (
	EnvLangValueParserRULE_dqstring    = 0
	EnvLangValueParserRULE_content     = 1
	EnvLangValueParserRULE_variable    = 2
	EnvLangValueParserRULE_escapedChar = 3
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4094) != 0 {
		{
			p.SetState(8)
			p.Content()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(EnvLangValueParserEOF)
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

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	EscapedChar() IEscapedCharContext
	TEXT() antlr.TerminalNode
	WS() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode

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

func (s *ContentContext) EscapedChar() IEscapedCharContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEscapedCharContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEscapedCharContext)
}

func (s *ContentContext) TEXT() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserTEXT, 0)
}

func (s *ContentContext) WS() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserWS, 0)
}

func (s *ContentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserNEWLINE, 0)
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
	p.EnterRule(localctx, 2, EnvLangValueParserRULE_content)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY, EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET, EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY, EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET, EnvLangValueParserSIMPLE_STRICT_VAR, EnvLangValueParserSIMPLE_VAR, EnvLangValueParserDOLLAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Variable()
		}

	case EnvLangValueParserESCAPED_CHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.EscapedChar()
		}

	case EnvLangValueParserTEXT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(18)
			p.Match(EnvLangValueParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EnvLangValueParserWS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(19)
			p.Match(EnvLangValueParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EnvLangValueParserNEWLINE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(20)
			p.Match(EnvLangValueParserNEWLINE)
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY() antlr.TerminalNode
	STRICT_VAR_WITH_ASSIGN_IF_UNSET() antlr.TerminalNode
	STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY() antlr.TerminalNode
	STRICT_VAR_WITH_DEFAULT_IF_UNSET() antlr.TerminalNode
	SIMPLE_STRICT_VAR() antlr.TerminalNode
	SIMPLE_VAR() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode

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

func (s *VariableContext) STRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY, 0)
}

func (s *VariableContext) STRICT_VAR_WITH_ASSIGN_IF_UNSET() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSTRICT_VAR_WITH_ASSIGN_IF_UNSET, 0)
}

func (s *VariableContext) STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY, 0)
}

func (s *VariableContext) STRICT_VAR_WITH_DEFAULT_IF_UNSET() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET, 0)
}

func (s *VariableContext) SIMPLE_STRICT_VAR() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSIMPLE_STRICT_VAR, 0)
}

func (s *VariableContext) SIMPLE_VAR() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSIMPLE_VAR, 0)
}

func (s *VariableContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserDOLLAR, 0)
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
	p.EnterRule(localctx, 4, EnvLangValueParserRULE_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254) != 0) {
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

// IEscapedCharContext is an interface to support dynamic dispatch.
type IEscapedCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCAPED_CHAR() antlr.TerminalNode

	// IsEscapedCharContext differentiates from other interfaces.
	IsEscapedCharContext()
}

type EscapedCharContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscapedCharContext() *EscapedCharContext {
	var p = new(EscapedCharContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_escapedChar
	return p
}

func InitEmptyEscapedCharContext(p *EscapedCharContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EnvLangValueParserRULE_escapedChar
}

func (*EscapedCharContext) IsEscapedCharContext() {}

func NewEscapedCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EscapedCharContext {
	var p = new(EscapedCharContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EnvLangValueParserRULE_escapedChar

	return p
}

func (s *EscapedCharContext) GetParser() antlr.Parser { return s.parser }

func (s *EscapedCharContext) ESCAPED_CHAR() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserESCAPED_CHAR, 0)
}

func (s *EscapedCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscapedCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EscapedCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.EnterEscapedChar(s)
	}
}

func (s *EscapedCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EnvLangValueListener); ok {
		listenerT.ExitEscapedChar(s)
	}
}

func (p *EnvLangValueParser) EscapedChar() (localctx IEscapedCharContext) {
	localctx = NewEscapedCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EnvLangValueParserRULE_escapedChar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(EnvLangValueParserESCAPED_CHAR)
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
