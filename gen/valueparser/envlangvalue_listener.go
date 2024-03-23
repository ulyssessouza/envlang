// Code generated from EnvLangValue.g4 by ANTLR 4.13.1. DO NOT EDIT.

package valueparser // EnvLangValue
import "github.com/antlr4-go/antlr/v4"

// EnvLangValueListener is a complete listener for a parse tree produced by EnvLangValueParser.
type EnvLangValueListener interface {
	antlr.ParseTreeListener

	// EnterDqstring is called when entering the dqstring production.
	EnterDqstring(c *DqstringContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// ExitDqstring is called when exiting the dqstring production.
	ExitDqstring(c *DqstringContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)
}
