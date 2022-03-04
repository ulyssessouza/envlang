// Code generated from EnvLangValue.g4 by ANTLR 4.9.3. DO NOT EDIT.

package valueparser // EnvLangValue
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEnvLangValueListener is a complete listener for a parse tree produced by EnvLangValueParser.
type BaseEnvLangValueListener struct{}

var _ EnvLangValueListener = &BaseEnvLangValueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEnvLangValueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEnvLangValueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEnvLangValueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEnvLangValueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDqstring is called when production dqstring is entered.
func (s *BaseEnvLangValueListener) EnterDqstring(ctx *DqstringContext) {}

// ExitDqstring is called when production dqstring is exited.
func (s *BaseEnvLangValueListener) ExitDqstring(ctx *DqstringContext) {}

// EnterStrictVar is called when production strictVar is entered.
func (s *BaseEnvLangValueListener) EnterStrictVar(ctx *StrictVarContext) {}

// ExitStrictVar is called when production strictVar is exited.
func (s *BaseEnvLangValueListener) ExitStrictVar(ctx *StrictVarContext) {}

// EnterSimpleVar is called when production simpleVar is entered.
func (s *BaseEnvLangValueListener) EnterSimpleVar(ctx *SimpleVarContext) {}

// ExitSimpleVar is called when production simpleVar is exited.
func (s *BaseEnvLangValueListener) ExitSimpleVar(ctx *SimpleVarContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseEnvLangValueListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseEnvLangValueListener) ExitVariable(ctx *VariableContext) {}

// EnterText is called when production text is entered.
func (s *BaseEnvLangValueListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseEnvLangValueListener) ExitText(ctx *TextContext) {}

// EnterSpace is called when production space is entered.
func (s *BaseEnvLangValueListener) EnterSpace(ctx *SpaceContext) {}

// ExitSpace is called when production space is exited.
func (s *BaseEnvLangValueListener) ExitSpace(ctx *SpaceContext) {}

// EnterDQEscape is called when production dQEscape is entered.
func (s *BaseEnvLangValueListener) EnterDQEscape(ctx *DQEscapeContext) {}

// ExitDQEscape is called when production dQEscape is exited.
func (s *BaseEnvLangValueListener) ExitDQEscape(ctx *DQEscapeContext) {}

// EnterSpecial is called when production special is entered.
func (s *BaseEnvLangValueListener) EnterSpecial(ctx *SpecialContext) {}

// ExitSpecial is called when production special is exited.
func (s *BaseEnvLangValueListener) ExitSpecial(ctx *SpecialContext) {}

// EnterContent is called when production content is entered.
func (s *BaseEnvLangValueListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseEnvLangValueListener) ExitContent(ctx *ContentContext) {}
