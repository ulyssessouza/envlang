package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ulyssessouza/envlang/parser"
)

type envVarListener struct {
	*parser.BaseEnvLangListener
}

func (l *envVarListener) ExitEntry(c *parser.EntryContext) {
	value := "NO_VALUE!"
	valueType := "NO_VALUE_TYPE!"

	id := c.Identifier().GetText()
	hasAssign := true
	if c.Assign() == nil || c.Assign().IsEmpty() {
		hasAssign = false
	}
	if hasAssign && c.Value() == nil {
		value = ""
	}

	if c.Value() != nil {
		value = c.Value().GetText()
		switch c.Value().GetOp().GetTokenType() {
		case parser.EnvLangParserTEXT:
			valueType = "TEXT"
		case parser.EnvLangParserDQSTRING:
			valueType = "DOUBLE_QUOTED"
		case parser.EnvLangParserSQSTRING:
			valueType = "SINGLE_QUOTED"
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.Value().GetOp().GetText()))
		}
	}

	fmt.Printf(
		"%q >>> %q >>> %q\n",
		valueType,
		id,
		value,
	)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream(`


A=aaa
B="bbb"
C =ccc
D= ddd
E='eee'
F=
G

H="my_value"


I = bar baz



J = "foo bar"

L="my
multi
line
entry"


N1=42
N2="42"`)

	// Create the Lexer
	lexer := parser.NewEnvLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewEnvLangParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&envVarListener{}, p.EnvFile())
}
