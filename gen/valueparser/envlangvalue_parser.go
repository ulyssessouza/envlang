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
		"", "STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY", "STRICT_VAR_WITH_DEFAULT_IF_UNSET",
		"SIMPLE_STRICT_VAR", "SIMPLE_VAR", "STR", "SPACE", "CRLF", "ANY",
	}
	staticData.RuleNames = []string{
		"dqstring", "content", "variable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 36, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 5, 0, 8, 8, 0, 10,
		0, 12, 0, 11, 9, 0, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0,
		5, 0, 20, 8, 0, 10, 0, 12, 0, 23, 9, 0, 1, 0, 3, 0, 26, 8, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 32, 8, 1, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 1,
		1, 0, 1, 4, 40, 0, 9, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0,
		6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 11, 1, 0, 0, 0, 9, 7, 1, 0, 0, 0,
		9, 10, 1, 0, 0, 0, 10, 25, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 12, 14, 5, 6,
		0, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 26, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 20, 5, 7, 0, 0,
		19, 18, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1,
		0, 0, 0, 22, 26, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 26, 5, 0, 0, 1, 25,
		15, 1, 0, 0, 0, 25, 21, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 1, 1, 0, 0,
		0, 27, 32, 3, 4, 2, 0, 28, 32, 5, 5, 0, 0, 29, 32, 5, 6, 0, 0, 30, 32,
		5, 7, 0, 0, 31, 27, 1, 0, 0, 0, 31, 28, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0,
		31, 30, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 34, 7, 0, 0, 0, 34, 5, 1, 0,
		0, 0, 5, 9, 15, 21, 25, 31,
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
	EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY = 1
	EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET          = 2
	EnvLangValueParserSIMPLE_STRICT_VAR                         = 3
	EnvLangValueParserSIMPLE_VAR                                = 4
	EnvLangValueParserSTR                                       = 5
	EnvLangValueParserSPACE                                     = 6
	EnvLangValueParserCRLF                                      = 7
	EnvLangValueParserANY                                       = 8
)

// EnvLangValueParser rules.
const (
	EnvLangValueParserRULE_dqstring = 0
	EnvLangValueParserRULE_content  = 1
	EnvLangValueParserRULE_variable = 2
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
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
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

func (s *DqstringContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(EnvLangValueParserSPACE)
}

func (s *DqstringContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSPACE, i)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(6)
				p.Content()
			}

		}
		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangValueParserSPACE {
			{
				p.SetState(12)
				p.Match(EnvLangValueParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(17)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangValueParserCRLF {
			{
				p.SetState(18)
				p.Match(EnvLangValueParserCRLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		{
			p.SetState(24)
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

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	STR() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	CRLF() antlr.TerminalNode

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

func (s *ContentContext) STR() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSTR, 0)
}

func (s *ContentContext) SPACE() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserSPACE, 0)
}

func (s *ContentContext) CRLF() antlr.TerminalNode {
	return s.GetToken(EnvLangValueParserCRLF, 0)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY, EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET, EnvLangValueParserSIMPLE_STRICT_VAR, EnvLangValueParserSIMPLE_VAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Variable()
		}

	case EnvLangValueParserSTR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Match(EnvLangValueParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EnvLangValueParserSPACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.Match(EnvLangValueParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EnvLangValueParserCRLF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.Match(EnvLangValueParserCRLF)
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

	// GetVar_ returns the var_ token.
	GetVar_() antlr.Token

	// SetVar_ sets the var_ token.
	SetVar_(antlr.Token)

	// Getter signatures
	STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY() antlr.TerminalNode
	STRICT_VAR_WITH_DEFAULT_IF_UNSET() antlr.TerminalNode
	SIMPLE_STRICT_VAR() antlr.TerminalNode
	SIMPLE_VAR() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	var_   antlr.Token
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

func (s *VariableContext) GetVar_() antlr.Token { return s.var_ }

func (s *VariableContext) SetVar_(v antlr.Token) { s.var_ = v }

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
		p.SetState(33)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableContext).var_ = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableContext).var_ = _ri
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
