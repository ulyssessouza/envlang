package handlers

import (
	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/valueparser"
)

const (
	pair = 2
)

var _ valueparser.EnvLangValueListener = &envLangValueListener{}

type envLangValueListener struct {
	valueparser.BaseEnvLangValueListener

	d      dao.EnvLangDao
	result string
}

func (l *envLangValueListener) ExitEveryRule(c antlr.ParserRuleContext) {
	log.Debugf("ExitEveryRule: %s", c.GetText())
}

func (l *envLangValueListener) EnterEveryRule(c antlr.ParserRuleContext) {
	log.Debugf("EnterEveryRule: %s", c.GetText())
}

func (l *envLangValueListener) append(s string) {
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
