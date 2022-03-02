// Code generated from EnvLangFile.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fileparser // EnvLangFile
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEnvLangFileListener is a complete listener for a parse tree produced by EnvLangFileParser.
type BaseEnvLangFileListener struct{}

var _ EnvLangFileListener = &BaseEnvLangFileListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEnvLangFileListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEnvLangFileListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEnvLangFileListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEnvLangFileListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEnvFile is called when production envFile is entered.
func (s *BaseEnvLangFileListener) EnterEnvFile(ctx *EnvFileContext) {}

// ExitEnvFile is called when production envFile is exited.
func (s *BaseEnvLangFileListener) ExitEnvFile(ctx *EnvFileContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseEnvLangFileListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseEnvLangFileListener) ExitEntry(ctx *EntryContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseEnvLangFileListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseEnvLangFileListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseEnvLangFileListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseEnvLangFileListener) ExitAssign(ctx *AssignContext) {}

// EnterValue is called when production value is entered.
func (s *BaseEnvLangFileListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseEnvLangFileListener) ExitValue(ctx *ValueContext) {}

// EnterNl is called when production nl is entered.
func (s *BaseEnvLangFileListener) EnterNl(ctx *NlContext) {}

// ExitNl is called when production nl is exited.
func (s *BaseEnvLangFileListener) ExitNl(ctx *NlContext) {}
