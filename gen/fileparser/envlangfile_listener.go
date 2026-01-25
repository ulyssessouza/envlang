// Code generated from EnvLangFile.g4 by ANTLR 4.13.2. DO NOT EDIT.

package fileparser // EnvLangFile
import "github.com/antlr4-go/antlr/v4"

// EnvLangFileListener is a complete listener for a parse tree produced by EnvLangFileParser.
type EnvLangFileListener interface {
	antlr.ParseTreeListener

	// EnterEnvFile is called when entering the envFile production.
	EnterEnvFile(c *EnvFileContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterInlineComment is called when entering the inlineComment production.
	EnterInlineComment(c *InlineCommentContext)

	// ExitEnvFile is called when exiting the envFile production.
	ExitEnvFile(c *EnvFileContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitInlineComment is called when exiting the inlineComment production.
	ExitInlineComment(c *InlineCommentContext)
}
