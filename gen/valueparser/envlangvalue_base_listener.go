// Code generated from EnvLangValue.g4 by ANTLR 4.13.1. DO NOT EDIT.

package valueparser // EnvLangValue
import "github.com/antlr4-go/antlr/v4"

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

// EnterVariable is called when production variable is entered.
func (s *BaseEnvLangValueListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseEnvLangValueListener) ExitVariable(ctx *VariableContext) {}

// EnterContent is called when production content is entered.
func (s *BaseEnvLangValueListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseEnvLangValueListener) ExitContent(ctx *ContentContext) {}
