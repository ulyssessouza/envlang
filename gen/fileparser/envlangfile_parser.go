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
		"", "'='", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "CRLF", "TEXT", "DQSTRING", "SQSTRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"envFile", "entry", "identifier", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 51, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5, 0,
		10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12,
		0, 20, 9, 0, 4, 0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 30, 8, 1, 1, 1, 5, 1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 1, 5,
		1, 39, 8, 1, 10, 1, 12, 1, 42, 9, 1, 1, 1, 3, 1, 45, 8, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 1, 1, 0, 3, 5, 54, 0, 11, 1,
		0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 48, 1, 0, 0, 0, 8, 10,
		5, 2, 0, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 11,
		12, 1, 0, 0, 0, 12, 21, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 14, 18, 3, 2, 1,
		0, 15, 17, 5, 2, 0, 0, 16, 15, 1, 0, 0, 0, 17, 20, 1, 0, 0, 0, 18, 16,
		1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		21, 14, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1,
		0, 0, 0, 24, 1, 1, 0, 0, 0, 25, 29, 3, 4, 2, 0, 26, 27, 5, 1, 0, 0, 27,
		30, 3, 6, 3, 0, 28, 30, 5, 1, 0, 0, 29, 26, 1, 0, 0, 0, 29, 28, 1, 0, 0,
		0, 29, 30, 1, 0, 0, 0, 30, 44, 1, 0, 0, 0, 31, 33, 5, 6, 0, 0, 32, 31,
		1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 40, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 39, 5, 2, 0, 0, 38, 37, 1,
		0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41,
		45, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 45, 5, 0, 0, 1, 44, 34, 1, 0, 0,
		0, 44, 43, 1, 0, 0, 0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 3, 0, 0, 47, 5, 1,
		0, 0, 0, 48, 49, 7, 0, 0, 0, 49, 7, 1, 0, 0, 0, 7, 11, 18, 23, 29, 34,
		40, 44,
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
	EnvLangFileParserASSIGN   = 1
	EnvLangFileParserCRLF     = 2
	EnvLangFileParserTEXT     = 3
	EnvLangFileParserDQSTRING = 4
	EnvLangFileParserSQSTRING = 5
	EnvLangFileParserSPACE    = 6
)

// EnvLangFileParser rules.
const (
	EnvLangFileParserRULE_envFile    = 0
	EnvLangFileParserRULE_entry      = 1
	EnvLangFileParserRULE_identifier = 2
	EnvLangFileParserRULE_value      = 3
)

// IEnvFileContext is an interface to support dynamic dispatch.
type IEnvFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode
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

func (s *EnvFileContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(EnvLangFileParserCRLF)
}

func (s *EnvFileContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserCRLF, i)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EnvLangFileParserCRLF {
		{
			p.SetState(8)
			p.Match(EnvLangFileParserCRLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EnvLangFileParserTEXT {
		{
			p.SetState(14)
			p.Entry()
		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangFileParserCRLF {
			{
				p.SetState(15)
				p.Match(EnvLangFileParserCRLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(23)
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
	ASSIGN() antlr.TerminalNode
	Value() IValueContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode

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

func (s *EntryContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserASSIGN, 0)
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

func (s *EntryContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(EnvLangFileParserSPACE)
}

func (s *EntryContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserSPACE, i)
}

func (s *EntryContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(EnvLangFileParserCRLF)
}

func (s *EntryContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(EnvLangFileParserCRLF, i)
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
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Identifier()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(26)
			p.Match(EnvLangFileParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Value()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(28)
			p.Match(EnvLangFileParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EnvLangFileParserSPACE {
			{
				p.SetState(31)
				p.Match(EnvLangFileParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(37)
					p.Match(EnvLangFileParserCRLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(43)
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
		p.SetState(46)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStr returns the str token.
	GetStr() antlr.Token

	// SetStr sets the str token.
	SetStr(antlr.Token)

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
	str    antlr.Token
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

func (s *ValueContext) GetStr() antlr.Token { return s.str }

func (s *ValueContext) SetStr(v antlr.Token) { s.str = v }

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
	p.EnterRule(localctx, 6, EnvLangFileParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ValueContext).str = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&56) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ValueContext).str = _ri
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
