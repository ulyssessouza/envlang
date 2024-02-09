package handlers

import (
	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/valueparser"
)

var _ valueparser.EnvLangValueListener = &envLangValueListener{}

type envLangValueListener struct {
	valueparser.BaseEnvLangValueListener

	d      dao.EnvLangDao
	result string
}

func (l *envLangValueListener) ExitVariable(c *valueparser.VariableContext) {
	log.Debugf("ExitVariable: %#v\n", c.GetText())
}

func (l *envLangValueListener) ExitStrictVar(c *valueparser.StrictVarContext) {
	varName := ""
	log.Debugf("ExitStrictVar: %#v\n", c.GetText())

	logVarName := "\tVarName: '"
	for _, t := range c.AllTEXT_NO_SPACE() {
		logVarName += t.GetText()
		varName += t.GetText()
	}
	log.Debugln(logVarName + "'")

	varVal, ok := l.d.Get(varName)
	if !ok {
		v := ""
		varVal = &v
	}

	l.result += *varVal
}

func (l *envLangValueListener) ExitSimpleVar(c *valueparser.SimpleVarContext) {
	varName := ""
	log.Debugf("ExitSimpleVar: %#v\n", c.GetText())

	logVarName := "\tVarName: '"
	for _, t := range c.AllTEXT_NO_SPACE() {
		logVarName += t.GetText()
		varName += t.GetText()
	}
	log.Debugln(logVarName + "'")

	varVal, ok := l.d.Get(varName)
	if !ok {
		v := ""
		varVal = &v
	}

	l.result += *varVal
}

func (l *envLangValueListener) ExitText(c *valueparser.TextContext) {
	log.Debugf("ExitText: %#v\n", c.GetText())
	l.Append(c.GetText())
}

func (l *envLangValueListener) ExitSpace(c *valueparser.SpaceContext) {
	log.Debugf("ExitSpace: %#v\n", c.GetText())
	l.Append(c.GetText())
}

func (l *envLangValueListener) ExitSpecial(c *valueparser.SpecialContext) {
	log.Debugf("ExitSpecial: %#v\n", c.GetText())
	l.Append(c.GetText())
}

func (l *envLangValueListener) Append(s string) {
	l.result += s
}

func newEnvLangValueListener(d dao.EnvLangDao) *envLangValueListener {
	return &envLangValueListener{
		d: d,
	}
}

func GetValue(d dao.EnvLangDao, s string) string {
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
