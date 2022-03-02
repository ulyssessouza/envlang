// Code generated from EnvLangValue.g4 by ANTLR 4.9.3. DO NOT EDIT.

package valueparser // EnvLangValue
import "github.com/antlr/antlr4/runtime/Go/antlr"

// EnvLangValueListener is a complete listener for a parse tree produced by EnvLangValueParser.
type EnvLangValueListener interface {
	antlr.ParseTreeListener

	// EnterDqstring is called when entering the dqstring production.
	EnterDqstring(c *DqstringContext)

	// EnterStrictVar is called when entering the strictVar production.
	EnterStrictVar(c *StrictVarContext)

	// EnterSimpleVar is called when entering the simpleVar production.
	EnterSimpleVar(c *SimpleVarContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterSpace is called when entering the space production.
	EnterSpace(c *SpaceContext)

	// EnterDQEscape is called when entering the dQEscape production.
	EnterDQEscape(c *DQEscapeContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// ExitDqstring is called when exiting the dqstring production.
	ExitDqstring(c *DqstringContext)

	// ExitStrictVar is called when exiting the strictVar production.
	ExitStrictVar(c *StrictVarContext)

	// ExitSimpleVar is called when exiting the simpleVar production.
	ExitSimpleVar(c *SimpleVarContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitSpace is called when exiting the space production.
	ExitSpace(c *SpaceContext)

	// ExitDQEscape is called when exiting the dQEscape production.
	ExitDQEscape(c *DQEscapeContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)
}
