package handlers

import (
	"fmt"

	antlr "github.com/antlr4-go/antlr/v4"

	"github.com/ulyssessouza/envlang/gen/valueparser"
	"github.com/ulyssessouza/envlang/logger"
	"github.com/ulyssessouza/envlang/store"
)

const (
	pair = 2
)

// ParameterExpansionError is raised when a parameter expansion error operator (? or :?) is triggered.
type ParameterExpansionError struct {
	VarName string
	Message string
}

func (e *ParameterExpansionError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.VarName, e.Message)
	}
	return fmt.Sprintf("%s: parameter not set", e.VarName)
}

var _ valueparser.EnvLangValueListener = &envLangValueListener{}

type envLangValueListener struct {
	valueparser.BaseEnvLangValueListener

	d      store.Store
	result string
}

func (l *envLangValueListener) ExitEveryRule(c antlr.ParserRuleContext) {
	logger.Debugf("ExitEveryRule: %s", c.GetText())
}

func (l *envLangValueListener) EnterEveryRule(c antlr.ParserRuleContext) {
	logger.Debugf("EnterEveryRule: %s", c.GetText())
}

func (l *envLangValueListener) append(s string) {
	l.result += s
}

func newEnvLangValueListener(d store.Store) *envLangValueListener {
	return &envLangValueListener{
		d: d,
	}
}

func GetValue(d store.Store, s string) string {
	is := antlr.NewInputStream(s)
	lexer := valueparser.NewEnvLangValueLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := valueparser.NewEnvLangValueParser(stream)
	parser.BuildParseTrees = true
	listener := newEnvLangValueListener(d)
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Dqstring())
	result := listener.result
	listener.result = ""

	return result
}
