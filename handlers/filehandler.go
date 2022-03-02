package handlers

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ulyssessouza/envlang/parsers/fileparser"
)

func NewEnvLangFileListener() *EnvLangFileListener {
	return &EnvLangFileListener{
		m: make(map[string]*string),
	}
}

var _ fileparser.EnvLangFileListener = &EnvLangFileListener{}

type EnvLangFileListener struct {
	*fileparser.BaseEnvLangFileListener

	m map[string]*string
}

func (l EnvLangFileListener) GetVariables() map[string]*string {
	return l.m
}

func (l *EnvLangFileListener) ExitEntry(c *fileparser.EntryContext) {
	var valuePtr *string = nil
	// valueType := "NO_VALUE_TYPE!"

	id := c.Identifier().GetText()
	hasAssign := true
	if c.Assign() == nil || c.Assign().IsEmpty() {
		hasAssign = false
	}
	if hasAssign && c.Value() == nil {
		v := ""
		valuePtr = &v
	}

	if c.Value() != nil {
		v := c.Value().GetText()
		valuePtr = &v
		switch c.Value().GetOp().GetTokenType() {
		case fileparser.EnvLangFileParserTEXT:
			// valueType = "TEXT         : "
		case fileparser.EnvLangFileParserSQSTRING:
			// valueType = "SINGLE_QUOTED: "
		case fileparser.EnvLangFileParserDQSTRING:
			// valueType = "DOUBLE_QUOTED: "
			v := getValue(v[1 : len(v)-1])
			valuePtr = &v
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.Value().GetOp().GetText()))
		}
	}

	id = strings.TrimSpace(id)
	l.m[id] = valuePtr

	// v := "NO_VALUE!"
	// if valuePtr != nil {
	// 	v = *valuePtr
	// }
	// fmt.Printf(
	// 	"%q >>> %q >>> %q\n",
	// 	valueType,
	// 	id,
	// 	v,
	// )
}

func GetVariables(s string) map[string]*string {
	is := antlr.NewInputStream(s)
	lexer := fileparser.NewEnvLangFileLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := fileparser.NewEnvLangFileParser(stream)
	listener := NewEnvLangFileListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.EnvFile())
	return listener.GetVariables()
}
