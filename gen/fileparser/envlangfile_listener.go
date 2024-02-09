// Code generated from EnvLangFile.g4 by ANTLR 4.13.1. DO NOT EDIT.

package fileparser // EnvLangFile
import "github.com/antlr4-go/antlr/v4"

// EnvLangFileListener is a complete listener for a parse tree produced by EnvLangFileParser.
type EnvLangFileListener interface {
	antlr.ParseTreeListener

	// EnterEnvFile is called when entering the envFile production.
	EnterEnvFile(c *EnvFileContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterNl is called when entering the nl production.
	EnterNl(c *NlContext)

	// ExitEnvFile is called when exiting the envFile production.
	ExitEnvFile(c *EnvFileContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitNl is called when exiting the nl production.
	ExitNl(c *NlContext)
}
